package notification

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Notification struct {
	ID         uuid.UUID `gorm:"primarykey;type:uuid;not null;" json:"id,omitempty"`
	DocumentID uuid.UUID `gorm:"type:uuid;index;not null;" json:"document_id,omitempty"`
	Date       time.Time `gorm:"type:time;not null;" json:"date,omitempty"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

// Prepare Notification object before inserting into database
func (n *Notification) Prepare() {
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
}

// Validate Notification object before inserting into database
func (n *Notification) Validate() error {
	if n.DocumentID == uuid.Nil {
		return errors.New("document_id is required")
	}
	if n.Date.IsZero() {
		return errors.New("date is required")
	}

	return nil
}

func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	// Create UUID ID
	n.ID = uuid.New()

	n.Prepare()

	err := n.Validate()
	if err != nil {
		return err
	}

	return nil
}

// TODO: Add InsertNotification function (and test it before creating grpc)

// Insert one Notification object into database
func (n *Notification) InsertOne(ctx context.Context, db *gorm.DB) error {
	res := db.WithContext(ctx).Table("notifications").Create(&n)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
