/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_TENOR_API_KEY: string;
  readonly VITE_BOT_VERSION: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
