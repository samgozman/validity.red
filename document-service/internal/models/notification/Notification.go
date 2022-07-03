package notification

import (
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
func (d *Notification) Prepare() {
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
}

// Validate Notification object before inserting into database
func (d *Notification) Validate() error {
	if d.DocumentID == uuid.Nil {
		return errors.New("document_id is required")
	}
	if d.Date.IsZero() {
		return errors.New("date is required")
	}

	return nil
}

func (d *Notification) BeforeCreate(tx *gorm.DB) error {
	// Create UUID ID
	d.ID = uuid.New()

	d.Prepare()

	err := d.Validate()
	if err != nil {
		return err
	}

	return nil
}
