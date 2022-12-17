package user

import (
	"context"
)

type UserRepository interface {
	InsertOne(ctx context.Context, user *User) error
	FindOne(ctx context.Context, query *User, fields string) (*User, error)
	Update(ctx context.Context, userID string, fields map[string]interface{}) error
}
