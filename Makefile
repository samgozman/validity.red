GATEWAY_BINARY=gatewayApp
USER_BINARY=userApp
LOGGER_BINARY=loggerApp
DOCUMENT_BINARY=documentApp

## starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## Build go binaries
build_go: build_gateway build_user build_document

## stops docker-compose (if running), builds all projects and starts docker compose
up_build: grpc_init_go grpc_init_rust build_go
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

# build proto files and copy them into Go services
grpc_init_go:
	@echo "Remove old gateway-service/proto folder"
	rm -r gateway-service/proto || true
	@echo "Create new gateway-service/proto folders"
	mkdir gateway-service/proto gateway-service/proto/user gateway-service/proto/logs gateway-service/proto/document gateway-service/proto/calendar
	@echo "Starting proto files generation for Go..."
	protoc --go_out=./user-service --go_opt=paths=source_relative --go-grpc_out=./user-service --go-grpc_opt=paths=source_relative proto/user.proto
	protoc --go_out=./document-service --go_opt=paths=source_relative --go-grpc_out=./document-service --go-grpc_opt=paths=source_relative proto/document.proto
	protoc --go_out=./gateway-service --go_opt=paths=source_relative --go-grpc_out=./gateway-service --go-grpc_opt=paths=source_relative proto/calendar.proto
	@echo "Copy pregenerated Go proto files into gateway-service"
	cp user-service/proto/* gateway-service/proto/user
	cp document-service/proto/* gateway-service/proto/document
	@echo "Move calendar files from default folder into nested"
	mv gateway-service/proto/calendar*.go gateway-service/proto/calendar
	@echo "Done!"

# To init grpc for rust services we need to copy proto files into rust services
# Build proccess will be done by cargo
grpc_init_rust:
	@echo "Remove old calendar-serviceproto folder"
	rm -r calendar-service/proto || true
	@echo "Create new calendar-service/proto folders"
	mkdir calendar-service/proto
	cp proto/calendar.proto calendar-service/proto
	@echo "Done!"

# Lint rust services
lint_rust:
	rustfmt calendar-service/src/*.rs --edition 2021

## builds the gateway binary as a linux executable
build_gateway:
	@echo "Building broker binary..."
	cd ./gateway-service && env GOOS=linux CGO_ENABLED=0 go build -o ${GATEWAY_BINARY} ./cmd/api
	@echo "Done!"

## builds the user binary as a linux executable
build_user:
	@echo "Building user binary..."
	cd ./user-service && env GOOS=linux CGO_ENABLED=0 go build -o ${USER_BINARY} ./cmd/api
	@echo "Done!"

## builds the document binary as a linux executable
build_document:
	@echo "Building document binary..."
	cd ./document-service && env GOOS=linux CGO_ENABLED=0 go build -o ${DOCUMENT_BINARY} ./cmd/api
	@echo "Done!"
