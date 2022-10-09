package user

import (
	"context"
)

type UserRepository interface {
	InsertOne(ctx context.Context, user *User) error
	FindOneByEmail(ctx context.Context, email string) (*User, error)
	GetCalendarId(ctx context.Context, userId string) (*User, error)
	Update(ctx context.Context, userId string, fields map[string]interface{}) error
}
