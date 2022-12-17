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

func (u *PostgresTestRepository) InsertOne(ctx context.Context, user *user.User) error {
	user.ID, _ = uuid.Parse("434377cf-7509-4cc0-9895-0afa683f0e56")
	return nil
}

func (u *PostgresTestRepository) FindOne(ctx context.Context, query *user.User, fields string) (*user.User, error) {
	userID, _ := uuid.Parse("434377cf-7509-4cc0-9895-0afa683f0e56")

	user := &user.User{
		ID:         userID,
		Email:      "me@example.com",
		Password:   "",
		IsVerified: true,
		CalendarID: "8gipfmoqt8mtucep",
		IVCalendar: make([]byte, user.IVCalendarLength),
		Timezone:   "Europe/London",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return user, nil
}

func (u *PostgresTestRepository) Update(ctx context.Context, userID string, fields map[string]interface{}) error {
	return nil
}
