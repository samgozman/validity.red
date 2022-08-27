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

type PostgresRepository struct {
	Conn *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		Conn: db,
	}
}

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
func (user *User) Prepare() {
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

// Validate User object before inserting into database
func (user *User) Validate() error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if len(user.Password) < 8 {
		return errors.New("password is too short, must be at least 8 characters")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
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

func (user *User) BeforeCreate(tx *gorm.DB) error {
	// Create UUID ID
	user.ID = uuid.New()

	user.Prepare()

	err := user.Validate()
	if err != nil {
		return err
	}

	return nil
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// Insert one User object into database
func (u *PostgresRepository) InsertOne(ctx context.Context, user *User) error {
	res := u.Conn.Table("users").Create(&user).WithContext(ctx)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// Find one user by email
func (u *PostgresRepository) FindOneByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	res := u.Conn.Table("users").First(&user, "email = ?", email).WithContext(ctx)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
