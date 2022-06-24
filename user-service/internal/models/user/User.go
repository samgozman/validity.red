package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string              `bson:"email,omitempty" json:"email,omitempty"`
	CreatedAt primitive.Timestamp `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt primitive.Timestamp `bson:"updated_at" json:"updated_at,omitempty"`
}

// Insert one User object into database
func InsertOne(ctx context.Context, db *mongo.Database, user User) error {
	_, err := db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
