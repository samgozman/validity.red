package main

import (
	"context"
	"os"

	"github.com/samgozman/validity.red/user/pkg/mongodb"

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

	client, ctx, cancel, err := mongodb.Connect("mongodb://users_mongodb/", credential)
	if err != nil {
		panic(err)
	}
	defer mongodb.Close(ctx, client, cancel)

	database := client.Database(dbname)

	// Create indexes for users email field
	// TODO: move it to "models" directory
	mod := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err = database.Collection("users").Indexes().CreateOne(ctx, mod)
	if err != nil {
		panic(err)
	}

	// Create app
	app := Config{
		ctx: ctx,
		db:  database,
	}

	// Start gRPC server
	// go app.gRPCListen()
	app.gRPCListen()
}
