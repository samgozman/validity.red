# Dependencies caching stage
# Generate recipe file for dependencies by cargo-chef
FROM rust:1.64-alpine3.16 AS planner
WORKDIR /usr/src/app
RUN apk add musl-dev
RUN cargo install cargo-chef
COPY . .
RUN cargo chef prepare --recipe-path recipe.json

# Build dependencies
FROM rust:1.64-alpine3.16 AS cacher
WORKDIR /usr/src/app
RUN apk add musl-dev libressl-dev
RUN cargo install cargo-chef
COPY --from=planner /usr/src/app/recipe.json recipe.json
RUN cargo chef cook --release --recipe-path recipe.json

# Build stage
FROM rust:1.64-alpine3.16 as builder
WORKDIR /usr/src/app
RUN apk add build-base protoc protobuf-dev libressl-dev
COPY --from=cacher /usr/src/app/target target
COPY --from=cacher /usr/local/cargo /usr/local/cargo
COPY . .
RUN cargo build --release

# Final stage
FROM gcr.io/distroless/cc-debian11
COPY --from=builder /usr/src/app/target/release/calendar-service /usr/local/bin/calendar-service
CMD ["calendar-service"]
