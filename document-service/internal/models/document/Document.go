package document

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/notification"
	proto "github.com/samgozman/validity.red/document/proto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DocumentDB struct {
	Conn *gorm.DB
}

func NewDocumentDB(db *gorm.DB) *DocumentDB {
	return &DocumentDB{
		Conn: db.Table("documents"),
	}
}

type Document struct {
	ID            uuid.UUID                   `gorm:"primarykey;type:uuid;not null;" json:"id,omitempty"`
	UserID        uuid.UUID                   `gorm:"type:uuid;index;not null;" json:"user_id,omitempty"`
	Type          proto.Type                  `gorm:"type:int;default:0" json:"type,omitempty"`
	Title         string                      `gorm:"size:100;not null;" json:"title,omitempty"`
	Description   string                      `gorm:"size:500;not null;" json:"description,omitempty"`
	ExpiresAt     time.Time                   `gorm:"default:0" json:"expires_at,omitempty"`
	Notifications []notification.Notification `gorm:"foreignKey:DocumentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID" json:"notifications,omitempty"`
	CreatedAt     time.Time                   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt     time.Time                   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

// Prepare Document object before inserting into database
func (d *Document) Prepare() {
	escapeCharacters := regexp.MustCompile(`(?m)<|>|\(|\)|;|\\|\/`)
	d.Title = escapeCharacters.ReplaceAllString(strings.TrimSpace(d.Title), "\\$0")
	d.Description = escapeCharacters.ReplaceAllString(strings.TrimSpace(d.Description), "\\$0")
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
	if len([]rune(d.Title)) > 100 {
		return errors.New("title length must be less than 100 characters")
	}
	if len([]rune(d.Description)) > 500 {
		return errors.New("description length must be less than 500 characters")
	}
	if d.ExpiresAt.IsZero() {
		return errors.New("expires_at is required")
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
func (db *DocumentDB) InsertOne(ctx context.Context, d *Document) error {
	res := db.Conn.WithContext(ctx).Create(&d)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (db *DocumentDB) UpdateOne(ctx context.Context, d *Document) error {
	res := db.Conn.
		WithContext(ctx).
		Where(&Document{ID: d.ID, UserID: d.UserID}).
		Updates(&Document{
			Type:        d.Type,
			Title:       d.Title,
			Description: d.Description,
			ExpiresAt:   d.ExpiresAt,
		})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("document not found or you don't have permission to update it")
	}

	return nil
}

// TODO: Implement "soft delete" feature
// TODO: Allow users to restore a document after deletion
// TODO: Delete documents with DeletedAt timestamp > 14d with CRON job
// @see: https://gorm.io/docs/delete.html#Soft-Delete
func (db *DocumentDB) DeleteOne(ctx context.Context, d *Document) error {
	res := db.Conn.
		WithContext(ctx).
		// Delete all associations also
		Select(clause.Associations).
		Delete(&Document{ID: d.ID, UserID: d.UserID})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("document not found or you don't have permission to delete it")
	}

	return nil
}

// Find one document
func (db *DocumentDB) FindOne(ctx context.Context, d *Document) error {
	res := db.Conn.
		WithContext(ctx).
		Model(&Document{}).
		Where(&Document{ID: d.ID, UserID: d.UserID}).
		First(&d)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("document not found or you don't have permission to view it")
	}

	return nil
}

// Checks if document is already exists in database
func (db *DocumentDB) Exists(ctx context.Context, d *Document) (bool, error) {
	var exist struct {
		Found bool
	}
	res := db.Conn.
		WithContext(ctx).
		Raw(
			"SELECT EXISTS(SELECT 1 FROM documents WHERE id = ? AND user_id = ?) as found",
			d.ID,
			d.UserID,
		).
		Scan(&exist)
	if res.Error != nil {
		return false, res.Error
	}

	return exist.Found, nil
}

// Find all documents by UserID
func (db *DocumentDB) FindAll(ctx context.Context, userId uuid.UUID) ([]Document, error) {
	var documents []Document

	res := db.Conn.
		WithContext(ctx).
		Model(&Document{}).
		Where(&Document{UserID: userId}).
		Find(&documents)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("documents not found")
	}

	return documents, nil
}
