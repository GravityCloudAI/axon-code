interface ImportMetaEnv {
  readonly VITE_axoncode_SERVER_HOST: string
  readonly VITE_axoncode_SERVER_PORT: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
