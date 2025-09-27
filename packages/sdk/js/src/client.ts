export * from "./gen/types.gen.js"
export { type Config as axoncodeClientConfig, axoncodeClient }

import { createClient } from "./gen/client/client.gen.js"
import { type Config } from "./gen/client/types.gen.js"
import { axoncodeClient } from "./gen/sdk.gen.js"

export function createaxoncodeClient(config?: Config) {
  const client = createClient(config)
  return new axoncodeClient({ client })
}
