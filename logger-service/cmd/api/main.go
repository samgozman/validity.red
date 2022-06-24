package main

import (
	"context"
	"os"

	"github.com/samgozman/validity.red/logger/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
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

	// Create index on field with TTL
	mod := mongo.IndexModel{
		Keys: bson.M{"created_at": 1},
		// Expire after 14 days
		Options: options.Index().SetExpireAfterSeconds(60 * 60 * 24 * 14),
	}
	_, err = database.Collection("logs").Indexes().CreateOne(ctx, mod)
	if err != nil {
		panic(err)
	}

	// Create app
	app := Config{
		ctx: ctx,
		db:  database,
	}

	app.gRPCListen()
}
