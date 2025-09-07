import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import path from "path";
import { getEntriesPattern } from "./utils/file-loader";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  build: {
    outDir: "../../src/static/ui",
    emptyOutDir: true,
    rollupOptions: {
      input: getEntriesPattern(),
      output: {
        entryFileNames: "[name].js",
        assetFileNames: "[name].[ext]",
      },
    },
  },
});
