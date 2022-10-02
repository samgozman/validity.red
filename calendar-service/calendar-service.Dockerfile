FROM rust:1.64-alpine3.16 as builder
WORKDIR /usr/src/app
COPY . .

RUN apk add build-base protoc protobuf-dev
RUN cargo build --release

FROM alpine:latest
COPY --from=builder /usr/src/app/target/release/calendar-service /usr/local/bin/app
CMD ["app"]
