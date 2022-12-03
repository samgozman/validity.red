import { fileURLToPath, URL } from "url";

import { defineConfig } from "vite";
import sentryVitePlugin from "@sentry/vite-plugin";
import vue from "@vitejs/plugin-vue";
import vueJsx from "@vitejs/plugin-vue-jsx";

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    sourcemap: true, // Source map generation must be turned on for Sentry
  },
  plugins: [
    vue(),
    vueJsx(),
    sentryVitePlugin({
      org: "samgozman",
      project: "validity-frontend-service",
      // Specify the directory containing build artifacts
      include: "./dist",
      // Auth tokens can be obtained from https://sentry.io/settings/account/api/auth-tokens/
      // and needs the `project:releases` and `org:read` scopes
      authToken: process.env.SENTRY_AUTH_TOKEN,
    }),
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
