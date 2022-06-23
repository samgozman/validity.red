BROKER_BINARY=brokerApp
AUTH_BINARY=authApp

## starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## stops docker-compose (if running), builds all projects and starts docker compose
up_build: grpc_init build_broker build_auth
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

grpc_init:
	@echo "Starting proto files generation..."
	protoc --go_out=./auth-service --go_opt=paths=source_relative --go-grpc_out=./auth-service --go-grpc_opt=paths=source_relative proto/auth.proto
	@echo "Done!"

## builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

## builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ./auth-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"