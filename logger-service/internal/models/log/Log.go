package log

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Log struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	Service   string              `bson:"service,omitempty" json:"service,omitempty"`   // Service name to which this log belongs
	LogLevel  string              `bson:"logLevel,omitempty" json:"logLevel,omitempty"` // INFO, DEBUG, WARN, ERROR, FATAL
	Message   string              `bson:"message,omitempty" json:"message,omitempty"`   // Message to make this log more helpful
	Error     any                 `bson:"error,omitempty" json:"error,omitempty"`
	CreatedAt primitive.Timestamp `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt primitive.Timestamp `bson:"updated_at" json:"updated_at,omitempty"`
}

// Insert one Log object into database
func InsertOne(ctx context.Context, db *mongo.Database, user Log) error {
	_, err := db.Collection("logs").InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
