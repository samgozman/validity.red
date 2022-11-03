package mocks

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/user/internal/models/user"
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
func (u *PostgresTestRepository) InsertOne(ctx context.Context, user *user.User) error {
	user.ID, _ = uuid.Parse("434377cf-7509-4cc0-9895-0afa683f0e56")
	return nil
}

func (u *PostgresTestRepository) FindOne(ctx context.Context, query []interface{}, fields string) (*user.User, error) {
	userId, _ := uuid.Parse("434377cf-7509-4cc0-9895-0afa683f0e56")
	user := &user.User{
		ID:          userId,
		Email:       "me@example.com",
		Password:    "",
		IsVerified:  true,
		CalendarID:  "8gipfmoqt8mtucep",
		IV_Calendar: make([]byte, 12),
		Timezone:    "Europe/London",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return user, nil
}

func (u *PostgresTestRepository) Update(ctx context.Context, userId string, fields map[string]interface{}) error {
	return nil
}
