package user

import (
	"context"
)

type UserRepository interface {
	InsertOne(ctx context.Context, user *User) error
	FindOneByEmail(ctx context.Context, email string) (*User, error)
}
