import { NoSuchModelError, type LanguageModel, type Provider as SDK } from "ai"
import path from "path"
import { mergeDeep, sortBy } from "remeda"
import z from "zod/v4"
import { BunProc } from "../bun"
import { Config } from "../config/config"
import { Global } from "../global"
import { Instance } from "../project/instance"
import { NamedError } from "../util/error"
import { Log } from "../util/log"
import { ModelsDev } from "./models"

export namespace Provider {
  const log = Log.create({ service: "provider" })

  type Source = "env" | "config" | "custom" | "api"

  const state = Instance.state(async () => {
    const config = await Config.get()
    const database = await ModelsDev.get()

    const providers: {
      [providerID: string]: {
        source: Source
        info: ModelsDev.Provider
        getModel?: (sdk: any, modelID: string) => Promise<any>
        options: Record<string, any>
      }
    } = {}
    const models = new Map<
      string,
      { providerID: string; modelID: string; info: ModelsDev.Model; language: LanguageModel; npm?: string }
    >()
    const sdk = new Map<number, SDK>()

    log.info("init")

    function mergeProvider(
      id: string,
      options: Record<string, any>,
      source: Source,
      getModel?: (sdk: any, modelID: string) => Promise<any>,
    ) {
      const provider = providers[id]
      if (!provider) {
        const info = database[id]
        if (!info) return
        if (info.api && !options["baseURL"]) options["baseURL"] = info.api
        providers[id] = {
          source,
          info,
          options,
          getModel,
        }
        return
      }
      provider.options = mergeDeep(provider.options, options)
      provider.source = source
      provider.getModel = getModel ?? provider.getModel
    }

    const configProviders = Object.entries(config.provider ?? {})

    for (const [providerID, provider] of configProviders) {
      const existing = database[providerID]
      const parsed: ModelsDev.Provider = {
        id: providerID,
        npm: provider.npm ?? existing?.npm,
        name: provider.name ?? existing?.name ?? providerID,
        env: provider.env ?? existing?.env ?? [],
        api: provider.api ?? existing?.api,
        models: existing?.models ?? {},
      }

      for (const [modelID, model] of Object.entries(provider.models ?? {})) {
        const existing = parsed.models[modelID]
        const parsedModel: ModelsDev.Model = {
          id: modelID,
          name: model.name ?? existing?.name ?? modelID,
          release_date: model.release_date ?? existing?.release_date,
          attachment: model.attachment ?? existing?.attachment ?? false,
          reasoning: model.reasoning ?? existing?.reasoning ?? false,
          temperature: model.temperature ?? existing?.temperature ?? false,
          tool_call: model.tool_call ?? existing?.tool_call ?? true,
          cost:
            !model.cost && !existing?.cost
              ? {
                  input: 0,
                  output: 0,
                  cache_read: 0,
                  cache_write: 0,
                }
              : {
                  cache_read: 0,
                  cache_write: 0,
                  ...existing?.cost,
                  ...model.cost,
                },
          options: {
            ...existing?.options,
            ...model.options,
          },
          limit: model.limit ??
            existing?.limit ?? {
              context: 0,
              output: 0,
            },
          provider: model.provider ?? existing?.provider,
        }
        parsed.models[modelID] = parsedModel
      }
      database[providerID] = parsed
    }

    // load config - ONLY load hardcoded provider from config
    for (const [providerID, provider] of configProviders) {
      mergeProvider(providerID, provider.options ?? {}, "config")
    }

    return {
      models,
      providers,
      sdk,
    }
  })

  export async function list() {
    return state().then((state) => state.providers)
  }

  async function getSDK(provider: ModelsDev.Provider, model: ModelsDev.Model) {
    return (async () => {
      using _ = log.time("getSDK", {
        providerID: provider.id,
      })
      const s = await state()
      const pkg = model.provider?.npm ?? provider.npm ?? provider.id
      const options = { ...s.providers[provider.id]?.options }
      if (pkg.includes("@ai-sdk/openai-compatible") && options["includeUsage"] === undefined) {
        options["includeUsage"] = true
      }
      const key = Bun.hash.xxHash32(JSON.stringify({ pkg, options }))
      const existing = s.sdk.get(key)
      if (existing) return existing
      const mod = await import(await BunProc.install(pkg, "latest"))
      if (options["timeout"] !== undefined) {
        // Only override fetch if user explicitly sets timeout
        options["fetch"] = async (input: any, init?: any) => {
          return await fetch(input, { ...init, timeout: options["timeout"] })
        }
      }
      const fn = mod[Object.keys(mod).find((key) => key.startsWith("create"))!]
      const loaded = fn({
        name: provider.id,
        ...options,
      })
      s.sdk.set(key, loaded)
      return loaded as SDK
    })().catch((e) => {
      throw new InitError({ providerID: provider.id }, { cause: e })
    })
  }

  export async function getProvider(providerID: string) {
    return state().then((s) => s.providers[providerID])
  }

  export async function getModel(providerID: string, modelID: string) {
    const key = `${providerID}/${modelID}`
    const s = await state()
    if (s.models.has(key)) return s.models.get(key)!

    log.info("getModel", {
      providerID,
      modelID,
    })

    const provider = s.providers[providerID]
    if (!provider) throw new ModelNotFoundError({ providerID, modelID })
    const info = provider.info.models[modelID]
    if (!info) throw new ModelNotFoundError({ providerID, modelID })
    const sdk = await getSDK(provider.info, info)

    try {
      const language = provider.getModel ? await provider.getModel(sdk, modelID) : sdk.languageModel(modelID)
      log.info("found", { providerID, modelID })
      s.models.set(key, {
        providerID,
        modelID,
        info,
        language,
        npm: info.provider?.npm ?? provider.info.npm,
      })
      return {
        modelID,
        providerID,
        info,
        language,
        npm: info.provider?.npm ?? provider.info.npm,
      }
    } catch (e) {
      if (e instanceof NoSuchModelError)
        throw new ModelNotFoundError(
          {
            modelID: modelID,
            providerID,
          },
          { cause: e },
        )
      throw e
    }
  }

  export async function getSmallModel(providerID: string) {
    const cfg = await Config.get()

    if (cfg.small_model) {
      const parsed = parseModel(cfg.small_model)
      return getModel(parsed.providerID, parsed.modelID)
    }

    const provider = await state().then((state) => state.providers[providerID])
    if (!provider) return
    const priority = ["3-5-haiku", "3.5-haiku", "gemini-2.5-flash", "gpt-5-nano"]
    for (const item of priority) {
      for (const model of Object.keys(provider.info.models)) {
        if (model.includes(item)) return getModel(providerID, model)
      }
    }
  }

  const priority = ["gemini-2.5-pro-preview", "gpt-5", "claude-sonnet-4"]
  export function sort(models: ModelsDev.Model[]) {
    return sortBy(
      models,
      [(model) => priority.findIndex((filter) => model.id.includes(filter)), "desc"],
      [(model) => (model.id.includes("latest") ? 0 : 1), "asc"],
      [(model) => model.id, "desc"],
    )
  }

  export async function defaultModel() {
    const cfg = await Config.get()
    log.info("defaultModel called", { cfg })
    if (cfg.model) return parseModel(cfg.model)

    // this will be adjusted when migration to opentui is complete,
    // for now we just read the tui state toml file directly
    //
    // NOTE: cannot just import file as toml without cleaning due to lack of
    // support for date/time references in Bun toml parser: https://github.com/oven-sh/bun/issues/22426
    const lastused = await Bun.file(path.join(Global.Path.state, "tui"))
      .text()
      .then((text) => {
        // remove the date/time references since Bun toml parser doesn't support yet
        const cleaned = text
          .split("\n")
          .filter((line) => !line.trim().startsWith("last_used ="))
          .join("\n")
        const state = Bun.TOML.parse(cleaned) as {
          recently_used_models?: {
            provider_id: string
            model_id: string
          }[]
        }
        const models = state?.recently_used_models ?? []
        if (models.length > 0) {
          return {
            providerID: models[0].provider_id,
            modelID: models[0].model_id,
          }
        }
      })
      .catch((error) => {
        log.error("failed to find last used model", {
          error,
        })
        return undefined
      })

    if (lastused) return lastused

    const provider = await list()
      .then((val) => Object.values(val))
      .then((x) => x.find((p) => !cfg.provider || Object.keys(cfg.provider).includes(p.info.id)))
    if (!provider) throw new Error("no providers found")
    const [model] = sort(Object.values(provider.info.models))
    if (!model) throw new Error("no models found")
    return {
      providerID: provider.info.id,
      modelID: model.id,
    }
  }

  export function parseModel(model: string) {
    const [providerID, ...rest] = model.split("/")
    return {
      providerID: providerID,
      modelID: rest.join("/"),
    }
  }

  export const ModelNotFoundError = NamedError.create(
    "ProviderModelNotFoundError",
    z.object({
      providerID: z.string(),
      modelID: z.string(),
    }),
  )

  export const InitError = NamedError.create(
    "ProviderInitError",
    z.object({
      providerID: z.string(),
    }),
  )
}
