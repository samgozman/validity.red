package document

import (
	"context"
	"errors"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/notification"
	"github.com/samgozman/validity.red/document/pkg/encryption"
	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	EncryptionKey = []byte(os.Getenv("ENCRYPTION_KEY"))
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
	ID             uuid.UUID                   `gorm:"primarykey;type:uuid;not null;" json:"id,omitempty"`
	UserID         uuid.UUID                   `gorm:"type:uuid;index;not null;" json:"user_id,omitempty"`
	Type           proto.Type                  `gorm:"type:int;default:0" json:"type,omitempty"`
	Title          string                      `gorm:"not null;" json:"title,omitempty"`
	Description    string                      `gorm:"" json:"description,omitempty"`
	IV_Title       []byte                      `gorm:"size:16;" json:"iv_title,omitempty"`
	IV_Description []byte                      `gorm:"size:16;" json:"iv_description,omitempty"`
	ExpiresAt      time.Time                   `gorm:"default:0" json:"expires_at,omitempty"`
	Notifications  []notification.Notification `gorm:"foreignKey:DocumentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:ID" json:"notifications,omitempty"`
	CreatedAt      time.Time                   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt      time.Time                   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

// Prepare Document object before inserting into database
func (d *Document) Prepare() {
	escapeCharacters := regexp.MustCompile(`(?m)<|>|\(|\)|;|\\|\/`)
	d.Title = escapeCharacters.ReplaceAllString(strings.TrimSpace(d.Title), "\\$0")
	d.Description = escapeCharacters.ReplaceAllString(strings.TrimSpace(d.Description), "\\$0")
}

// Validate Document object before inserting into database
func (d *Document) Validate() error {
	if d.UserID == uuid.Nil {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	if d.Title == "" {
		return status.Error(codes.InvalidArgument, "title is required")
	}
	if len([]rune(d.Title)) > 100 {
		return status.Error(codes.InvalidArgument, "title length must be less than 100 characters")
	}
	if len([]rune(d.Description)) > 500 {
		return status.Error(codes.InvalidArgument, "description length must be less than 500 characters")
	}
	if d.ExpiresAt.IsZero() {
		return status.Error(codes.InvalidArgument, "expires_at is required")
	}

	return nil
}

// Encrypt document title and description
func (d *Document) Encrypt() error {
	iv_title, err := encryption.GenerateRandomIV(encryption.BlockSize)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	// TODO: Do not encrypt description if it is empty
	iv_description, err := encryption.GenerateRandomIV(encryption.BlockSize)
	if err != nil {
		sentry.CaptureException(err)
		return status.Error(codes.Internal, err.Error())
	}

	encryptedTitle, err := encryption.EncryptAES(EncryptionKey, iv_title, d.Title)
	if err != nil {
		sentry.CaptureException(err)
		return status.Error(codes.Internal, err.Error())
	}

	encryptedDesc, err := encryption.EncryptAES(EncryptionKey, iv_description, d.Description)
	if err != nil {
		sentry.CaptureException(err)
		return status.Error(codes.Internal, err.Error())
	}

	d.Title = string(encryptedTitle)
	d.Description = string(encryptedDesc)
	d.IV_Title = iv_title
	d.IV_Description = iv_description

	return nil
}

func (d *Document) Decrypt() error {
	if d.IV_Title != nil {
		title, err := encryption.DecryptAES(EncryptionKey, d.IV_Title, d.Title)
		if err != nil {
			sentry.CaptureException(err)
			return status.Error(codes.Internal, err.Error())
		}
		d.Title = string(title)
	}

	if d.IV_Description != nil {
		desc, err := encryption.DecryptAES(EncryptionKey, d.IV_Description, d.Description)
		if err != nil {
			sentry.CaptureException(err)
			return status.Error(codes.Internal, err.Error())
		}
		d.Description = string(desc)
	}

	return nil
}

func (d *Document) BeforeCreate(tx *gorm.DB) error {
	// Create UUID ID
	d.ID = uuid.New()

	d.Prepare()

	err := d.Validate()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	err = d.Encrypt()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (d *Document) BeforeUpdate(tx *gorm.DB) error {
	d.Prepare()

	// TODO: Add validation for update event

	err := d.Encrypt()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (d *Document) AfterFind(tx *gorm.DB) error {
	err := d.Decrypt()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

// Insert one Document object into database
func (db *DocumentDB) InsertOne(ctx context.Context, d *Document) error {
	res := db.Conn.WithContext(ctx).Create(&d)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrInvalidData) ||
			errors.Is(res.Error, gorm.ErrInvalidValue) ||
			errors.Is(res.Error, gorm.ErrInvalidValueOfLength) {
			return status.Error(codes.InvalidArgument, "invalid document data")
		}
		sentry.CaptureException(res.Error)
		return status.Error(codes.Internal, res.Error.Error())
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
		if errors.Is(res.Error, gorm.ErrInvalidData) ||
			errors.Is(res.Error, gorm.ErrInvalidValue) ||
			errors.Is(res.Error, gorm.ErrInvalidValueOfLength) {
			return status.Error(codes.InvalidArgument, "invalid document data")
		}
		sentry.CaptureException(res.Error)
		return status.Error(codes.Internal, res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return status.Error(codes.NotFound, "document not found")
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
		sentry.CaptureException(res.Error)
		return status.Error(codes.Internal, res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return status.Error(codes.NotFound, "document not found")
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
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return status.Error(codes.NotFound, "document not found")
		}
		sentry.CaptureException(res.Error)
		return status.Error(codes.Internal, res.Error.Error())
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
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, status.Error(codes.NotFound, "document not found")
		}
		sentry.CaptureException(res.Error)
		return false, status.Error(codes.Internal, res.Error.Error())
	}

	return exist.Found, nil
}

// Find all documents by UserID
func (db *DocumentDB) FindAll(ctx context.Context, userId uuid.UUID) ([]Document, error) {
	var documents = []Document{}

	res := db.Conn.
		WithContext(ctx).
		Model(&Document{}).
		Where(&Document{UserID: userId}).
		Find(&documents)

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	return documents, nil
}

// Count user documents
func (db *DocumentDB) Count(ctx context.Context, userId uuid.UUID) (int64, error) {
	var count int64

	res := db.Conn.
		WithContext(ctx).
		Model(&Document{}).
		Where(&Document{UserID: userId}).
		Count(&count)

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return 0, status.Error(codes.Internal, res.Error.Error())
	}

	return count, nil
}

// Get count for all used document types
func (db *DocumentDB) CountTypes(ctx context.Context, userId uuid.UUID) ([]*proto.DocumentTypesCount, error) {
	var types = []*proto.DocumentTypesCount{}
	res := db.Conn.
		WithContext(ctx).
		Raw(
			"SELECT type, COUNT(*) FROM documents WHERE user_id = ? GROUP BY type",
			userId,
		).
		Scan(&types)

	if res.Error != nil {
		sentry.CaptureException(res.Error)
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	return types, nil
}

// Find top N latest documents sorted by expiration date
func (db *DocumentDB) FindLatest(ctx context.Context, userId uuid.UUID, limit int) ([]Document, error) {
	var documents = []Document{}

	// TODO: Specify attributes to fetch
	res := db.Conn.
		WithContext(ctx).
		Select("id, type, title, expires_at, iv_title").
		Model(&Document{}).
		Where(&Document{UserID: userId}).
		Order("expires_at ASC").
		Limit(limit).
		Find(&documents)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "document not found")
		}
		sentry.CaptureException(res.Error)
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	return documents, nil
}
