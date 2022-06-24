// ! Move this code into a tool box for this project
package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// This method closes mongoDB connection and   context.
func Close(ctx context.Context, client *mongo.Client, cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

// Connect to the mongo database
func Connect(uri string, auth options.Credential) (*mongo.Client, context.Context, context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 120 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(auth))
	return client, ctx, cancel, err
}
