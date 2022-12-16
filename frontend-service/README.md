# frontend-service

This service is the frontend SPA for [Validity.Red](https://validity.red), written in [Vue 3](https://v3.vuejs.org/) with [TypeScript](https://www.typescriptlang.org/). In production, it is served by [Nginx](https://www.nginx.com/) as a static site.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) with the following plugins:

- [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar)
- [TypeScript Vue Plugin](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin)
- [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)
- [Code Spell Checker](https://marketplace.visualstudio.com/items?itemName=streetsidesoftware.code-spell-checker)
- [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)
- [Tailwind CSS IntelliSense](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss)
- [Better Comments](https://marketplace.visualstudio.com/items?itemName=aaron-bond.better-comments)

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

This will normally be done by the [CI/CD pipeline](https://github.com/samgozman/validity.red/blob/main/.github/workflows/deploy_spa.yml), but you can run it locally if you want.

For that, you will need to set the `SENTRY_AUTH_TOKEN` environment variable which is used by `sentryVitePlugin` to upload source maps. You can get the token from [here](https://sentry.io/settings/account/api/auth-tokens/).

If you do not wish to create a Sentry project just for a local production build, you can remove the `sentryVitePlugin` from `vite.config.ts`.

```sh
npm run build
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```
