export namespace Flag {
  export const axoncode_AUTO_SHARE = truthy("axoncode_AUTO_SHARE")
  export const axoncode_CONFIG = process.env["axoncode_CONFIG"]
  export const axoncode_CONFIG_CONTENT = process.env["axoncode_CONFIG_CONTENT"]
  export const axoncode_DISABLE_AUTOUPDATE = truthy("axoncode_DISABLE_AUTOUPDATE")
  export const axoncode_DISABLE_PRUNE = truthy("axoncode_DISABLE_PRUNE")
  export const axoncode_PERMISSION = process.env["axoncode_PERMISSION"]
  export const axoncode_DISABLE_DEFAULT_PLUGINS = truthy("axoncode_DISABLE_DEFAULT_PLUGINS")
  export const axoncode_DISABLE_LSP_DOWNLOAD = truthy("axoncode_DISABLE_LSP_DOWNLOAD")
  export const axoncode_ENABLE_EXPERIMENTAL_MODELS = truthy("axoncode_ENABLE_EXPERIMENTAL_MODELS")
  export const axoncode_DISABLE_AUTOCOMPACT = truthy("axoncode_DISABLE_AUTOCOMPACT")

  // Experimental
  export const axoncode_EXPERIMENTAL_WATCHER = truthy("axoncode_EXPERIMENTAL_WATCHER")

  function truthy(key: string) {
    const value = process.env[key]?.toLowerCase()
    return value === "true" || value === "1"
  }
}
