import { createApp } from "vue";
import { init as sentryInit, vueRouterInstrumentation } from "@sentry/vue";
import { BrowserTracing } from "@sentry/browser";
import App from "./App.vue";
import router from "./router";

import "./assets/main.css";

const app = createApp(App);

sentryInit({
  app,
  dsn: import.meta.env.VITE_SENTRY_DSN,
  integrations: [
    new BrowserTracing({
      routingInstrumentation: vueRouterInstrumentation(router),
      tracePropagationTargets: [
        "localhost",
        "validity.red",
        "validity.host.extr.app",
        /^\//,
      ],
    }),
  ],
  tracesSampleRate: 0.2,
  sampleRate: 1.0,
  ignoreErrors: [
    // Random plugins/extensions
    "top.GLOBALS",
  ],
  denyUrls: [
    // Chrome extensions
    /extensions\//i,
    /^chrome:\/\//i,
  ],
});

app.use(router);

app.mount("#app");
