package user

import (
	"context"
	"errors"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// Id will be set as primaryKey by default
	ID        uuid.UUID `gorm:"type:uuid" json:"id,omitempty"`
	Email     string    `gorm:"uniqueIndex" json:"email,omitempty"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

// Prepare User object before inserting into database
func (u *User) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validate User object before inserting into database
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email required")
	}
	// TODO: Validate email to be of valid format

	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// Create UUID ID
	u.ID = uuid.New()

	u.Prepare()

	err := u.Validate()
	if err != nil {
		return err
	}

	return nil
}

// Insert one User object into database
func InsertOne(ctx context.Context, db *gorm.DB, u *User) error {
	res := db.Table("users").Create(&u).WithContext(ctx)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
