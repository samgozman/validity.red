package user

import (
	"context"
	"crypto/rand"
	"errors"
	"html"
	"math/big"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	ID          uuid.UUID `gorm:"type:uuid" json:"id,omitempty"`
	Email       string    `gorm:"uniqueIndex;size:100;not null;" json:"email,omitempty"`
	Password    string    `gorm:"not null;" json:"password"`
	IsVerified  bool      `gorm:"type:bool;default:false;not null;" json:"is_verified"`
	CalendarID  string    `gorm:"uniqueIndex;size:32;" json:"calendar_id,omitempty"`
	IV_Calendar []byte    `gorm:"size:12;" json:"iv_calendar,omitempty"`
	Timezone    string    `gorm:"size:50;default:Etc/UTC" json:"timezone,omitempty"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

// Prepare User object before inserting into database
func (user *User) Prepare() {
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.CalendarID = GenerateRandomString(32)
}

// Validate User object before inserting into database
func (user *User) Validate() error {
	if user.Email == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if user.Password == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	if len(user.Password) < 8 {
		return status.Error(codes.InvalidArgument, "password is too short, must be at least 8 characters")
	}
	if len(user.Password) > 64 {
		return status.Error(codes.InvalidArgument, "password is too long, must be between 8 - 64 characters")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return status.Error(codes.InvalidArgument, "invalid email")
	}

	return nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateRandomString(n int) string {
	var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		str[i] = chars[num.Int64()]
	}
	return string(str)
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
		if errors.Is(res.Error, gorm.ErrInvalidData) ||
			errors.Is(res.Error, gorm.ErrInvalidValue) ||
			errors.Is(res.Error, gorm.ErrInvalidValueOfLength) {
			return status.Error(codes.InvalidArgument, "invalid user data")
		}

		if strings.Contains(res.Error.Error(), "SQLSTATE 23505") {
			return status.Error(codes.AlreadyExists, "user is already exists")
		}

		return status.Error(codes.Internal, res.Error.Error())
	}

	return nil
}

// Find one user by "query" with selected "fields".
//
// Example:
//
// query: User{Email: "example@email.com"}
//
// fields: "calendar_id, iv_calendar, timezone"
func (u *PostgresRepository) FindOne(ctx context.Context, query *User, fields string) (*User, error) {
	user := &User{}
	res := u.Conn.Table("users").
		Select(fields).
		Where(query).
		First(&user).
		WithContext(ctx)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	return user, nil
}

func (u *PostgresRepository) Update(ctx context.Context, userId string, fields map[string]interface{}) error {
	res := u.Conn.WithContext(ctx).
		Table("users").
		Model(&User{}).
		Where("id = ?", userId).
		Updates(fields)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrInvalidData) ||
			errors.Is(res.Error, gorm.ErrInvalidValue) ||
			errors.Is(res.Error, gorm.ErrInvalidValueOfLength) {
			return status.Error(codes.InvalidArgument, "invalid user data")
		}
		return status.Error(codes.Internal, res.Error.Error())
	}
	if res.RowsAffected == 0 {
		return status.Error(codes.NotFound, "user not found")
	}

	return nil
}
