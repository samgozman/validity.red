# calendar-service

This service is the calendar service for [Validity.Red](https://validity.red), written in [Rust](https://www.rust-lang.org/).
It is used to generate and serve .ics calendar files with encrypted data.

The calendar service is connected only to the gateway service via gRPC and Protobuf.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) with the following plugins:

- [rust-analyzer](https://marketplace.visualstudio.com/items?itemName=rust-lang.rust-analyzer)
- [crates](https://marketplace.visualstudio.com/items?itemName=serayuzgur.crates)
- [Code Spell Checker](https://marketplace.visualstudio.com/items?itemName=streetsidesoftware.code-spell-checker)
- [Better Comments](https://marketplace.visualstudio.com/items?itemName=aaron-bond.better-comments)

## Project Setup

From the project root, run the following command to prepare protos:

```sh
make grpc_init_rust
```

and then run:

```sh
make build_rust
```

### Run for Development

This project is running in the docker-compose as a part of the project.
You can the project with the following command from the project root:

```sh
make up_build
```

This will build Rust and Go projects and run it in the docker-compose with all the necessary services.

### Unusual build features you should be aware of

As far as today, rust build time is not great at all.
It is not uncommon to wait for 10-15 minutes for a build to finish in a simple project.
To speed up the build process, I have added the [cargo-chef](https://github.com/LukeMathWalker/cargo-chef) tool.

This tool allows you to cache the build dependencies and reuse them in the next build by utilizing the docker layer caching.
It is done by splitting the build process into 4 steps which you can find in the `calendar-service.Dockerfile`.

This comes very handy when you are developing a project and you want to rebuild it often. So you need to wait
a little longer on the first build, but then you can rebuild it in a few seconds.

The downside of this approach, besides the added complexity, is that I rely on the not yet stable & reliable
tool which is not used by the majority of the community. But I'd like you to give it a try!

### Lint with `rustfmt`

From the project root, run:

```sh
make lint_rust
```
