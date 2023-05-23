# gateway-service

This service is the gateway service for [Validity](https://validity.extr.app), written in [Go](https://go.dev).
It is used to provide a REST API for the frontend service.

The gateway service is connected to the user service, document service and calendar service via gRPC and Protobuf.
Besides that, it is connected to the Redis database.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) with the following plugins:

- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [Go Test Explorer](https://marketplace.visualstudio.com/items?itemName=premparihar.gotestexplorer)
- [vscode-proto3](https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3)
- [Code Spell Checker](https://marketplace.visualstudio.com/items?itemName=streetsidesoftware.code-spell-checker)
- [Better Comments](https://marketplace.visualstudio.com/items?itemName=aaron-bond.better-comments)

With the following settings for Go linter
[https://golangci-lint.run/usage/integrations/#go-for-visual-studio-code](https://golangci-lint.run/usage/integrations/#go-for-visual-studio-code)

## Project Setup

From the project root, run the following command to prepare protos:

```sh
make grpc_init_go
```

and then run, if you want to create binary:

```sh
make build_document
```

### Run for Development

This service is running in the docker-compose as a part of the project.
You can run the project with the following command from the project root:

```sh
make up_build
```

This will build Rust and Go projects and run it in the docker-compose with all the necessary services.

The gateway server exposes port `8080`.

### Lint with [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)

From the project root, run:

```sh
make lint_go
```

*Linter should be installed locally.*
