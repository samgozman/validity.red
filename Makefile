BROKER_BINARY=brokerApp
USER_BINARY=userApp
LOGGER_BINARY=loggerApp
DOCUMENT_BINARY=documentApp

## starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## stops docker-compose (if running), builds all projects and starts docker compose
up_build: grpc_init build_broker build_user build_logger build_document
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

# build proto files and copy them into services
grpc_init:
	@echo "Starting proto files generation..."
	protoc --go_out=./user-service --go_opt=paths=source_relative --go-grpc_out=./user-service --go-grpc_opt=paths=source_relative proto/user.proto
	protoc --go_out=./logger-service --go_opt=paths=source_relative --go-grpc_out=./logger-service --go-grpc_opt=paths=source_relative proto/logs.proto
	protoc --go_out=./document-service --go_opt=paths=source_relative --go-grpc_out=./document-service --go-grpc_opt=paths=source_relative proto/document.proto
	@echo "Remove old broker-service/proto folder"
	rm -r broker-service/proto || true
	@echo "Copy pregenerated proto files into broker-service"
	mkdir broker-service/proto broker-service/proto/user broker-service/proto/logs broker-service/proto/document
	cp user-service/proto/* broker-service/proto/user
	cp logger-service/proto/* broker-service/proto/logs
	cp document-service/proto/* broker-service/proto/document
	@echo "Done!"

## builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
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

## builds the logger binary as a linux executable
build_logger:
	@echo "Building logger binary..."
	cd ./logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "Done!"