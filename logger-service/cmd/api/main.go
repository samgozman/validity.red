package main

import (
	"context"
	"os"

	"github.com/samgozman/validity.red/logger/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	ctx context.Context
	db  *mongo.Database
}

func main() {
	// Start mongodb server
	dbname := os.Getenv("MONGODB_NAME")

	credential := options.Credential{
		Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	}

	client, ctx, cancel, err := mongodb.Connect("mongodb://logger_mongodb/", credential)
	if err != nil {
		panic(err)
	}
	defer mongodb.Close(ctx, client, cancel)

	database := client.Database(dbname)

	// Create app
	app := Config{
		ctx: ctx,
		db:  database,
	}

	app.gRPCListen()
}
