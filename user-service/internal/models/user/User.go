package user

import (
	"context"
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	// Id will be set as primaryKey by default
	ID         uuid.UUID `gorm:"type:uuid" json:"id,omitempty"`
	Email      string    `gorm:"uniqueIndex;size:100;not null;" json:"email,omitempty"`
	Password   string    `gorm:"size:100;not null;" json:"password"`
	IsVerified bool      `gorm:"type:bool;default:false;not null;" json:"is_verified"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
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
		return errors.New("email is required")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid email")
	}

	return nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
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

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Insert one User object into database
func (u *User) InsertOne(ctx context.Context, db *gorm.DB) error {
	res := db.Table("users").Create(&u).WithContext(ctx)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// Find one user by email
func (u *User) FindOneByEmail(ctx context.Context, db *gorm.DB) (*User, error) {
	res := db.Table("users").First(&u, "email = ?", u.Email).WithContext(ctx)
	if res.Error != nil {
		return nil, res.Error
	}

	return u, nil
}
