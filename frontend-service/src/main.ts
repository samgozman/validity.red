import { createApp } from "vue";
import * as Sentry from "@sentry/vue";
import { BrowserTracing } from "@sentry/tracing";
import App from "./App.vue";
import router from "./router";

import "./assets/main.css";

const app = createApp(App);

Sentry.init({
  app,
  dsn: "https://c1411a0bdc4d4bdd922150dccf5a0df2@o1070792.ingest.sentry.io/4504265281437696",
  integrations: [
    new BrowserTracing({
      routingInstrumentation: Sentry.vueRouterInstrumentation(router),
      tracePropagationTargets: ["localhost", "validity.red", /^\//],
    }),
  ],
  tracesSampleRate: 0.2,
  sampleRate: 1.0,
});

app.use(router);

app.mount("#app");
