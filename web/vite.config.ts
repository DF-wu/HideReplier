import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import { fileURLToPath, URL } from "node:url";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: [
      {
        // Ship highlight.js core only; see src/shims/highlight.ts. The regex
        // must match the bare specifier only so the shim's own
        // "highlight.js/lib/core" import still resolves to the package.
        find: /^highlight\.js$/,
        replacement: fileURLToPath(
          new URL("./src/shims/highlight.ts", import.meta.url)
        ),
      },
    ],
  },
  build: {
    target: "es2020",
    sourcemap: false,
    // The server precompresses assets at image build time; skipping the
    // gzip size report shaves a little off every CI build.
    reportCompressedSize: false,
    chunkSizeWarningLimit: 600,
  },
});
