package document

import (
	"context"
	"errors"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Document struct {
	ID            uuid.UUID   `gorm:"type:uuid" json:"id,omitempty"`
	UserID        uuid.UUID   `gorm:"type:uuid;uniqueIndex;not null;" json:"user_id,omitempty"`
	Type          uint8       `gorm:"type:tinyint;default:0" json:"type,omitempty"` // Type number defined in proto Enum
	Title         string      `gorm:"size:100;not null;" json:"title,omitempty"`
	Description   string      `gorm:"size:500;not null;" json:"description,omitempty"`
	ExpiresAt     time.Time   `gorm:"default:0" json:"expires_at,omitempty"`
	Notifications []time.Time `gorm:"type:time[];default:[]" json:"notifications,omitempty"`
	CreatedAt     time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt     time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

// Prepare Document object before inserting into database
func (d *Document) Prepare() {
	d.Title = html.EscapeString(strings.TrimSpace(d.Title))
	d.Description = html.EscapeString(strings.TrimSpace(d.Description))
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
}

// Validate Document object before inserting into database
func (d *Document) Validate() error {
	if d.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	if d.Title == "" {
		return errors.New("title is required")
	}
	if d.Description == "" {
		return errors.New("description is required")
	}
	if len([]rune(d.Description)) > 500 {
		return errors.New("description length must be less than 500 characters")
	}
	if d.ExpiresAt.IsZero() {
		return errors.New("expiresAt is required")
	}

	return nil
}

func (d *Document) BeforeCreate(tx *gorm.DB) error {
	// Create UUID ID
	d.ID = uuid.New()

	d.Prepare()

	err := d.Validate()
	if err != nil {
		return err
	}

	return nil
}

// Insert one Document object into database
func InsertOne(ctx context.Context, db *gorm.DB, d *Document) error {
	res := db.Table("documents").Create(&d).WithContext(ctx)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
