package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresTestRepository struct {
	Conn *gorm.DB
}

func NewPostgresTestRepository(db *gorm.DB) *PostgresTestRepository {
	return &PostgresTestRepository{
		Conn: db,
	}
}

// Insert one User object into database
func (u *PostgresTestRepository) InsertOne(ctx context.Context, user *User) error {
	user.ID, _ = uuid.Parse("434377cf-7509-4cc0-9895-0afa683f0e56")
	return nil
}

// Find one user by email
func (u *PostgresTestRepository) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	userId, _ := uuid.Parse("434377cf-7509-4cc0-9895-0afa683f0e56")
	user := &User{
		ID:         userId,
		Email:      "me@example.com",
		Password:   "",
		IsVerified: true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return user, nil
}
