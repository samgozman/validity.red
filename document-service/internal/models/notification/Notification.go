package notification

import (
	"context"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type NotificationDB struct {
	Conn *gorm.DB
}

func NewNotificationDB(db *gorm.DB) *NotificationDB {
	return &NotificationDB{
		Conn: db.Table("notifications"),
	}
}

type Notification struct {
	ID         uuid.UUID `gorm:"primarykey;type:uuid;not null;" json:"id,omitempty"`
	UserID     uuid.UUID `gorm:"type:uuid;index;not null;" json:"user_id,omitempty"`
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
	if n.UserID == uuid.Nil {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	if n.DocumentID == uuid.Nil {
		return status.Error(codes.InvalidArgument, "document_id is required")
	}
	if n.Date.IsZero() {
		return status.Error(codes.InvalidArgument, "date is required")
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

// Insert one Notification object into database
func (db *NotificationDB) InsertOne(ctx context.Context, n *Notification) error {
	res := db.Conn.WithContext(ctx).Create(&n)
	if res.Error != nil {
		return status.Error(codes.Internal, res.Error.Error())
	}

	return nil
}

func (db *NotificationDB) DeleteOne(ctx context.Context, n *Notification) error {
	res := db.Conn.
		WithContext(ctx).
		Where(&Notification{ID: n.ID, DocumentID: n.DocumentID}).
		Delete(&Notification{})

	if res.Error != nil {
		return status.Error(codes.Internal, res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return status.Error(codes.NotFound, "notification not found")
	}

	return nil
}

func (db *NotificationDB) FindAll(ctx context.Context, documentID uuid.UUID) ([]Notification, error) {
	var notifications []Notification

	res := db.Conn.
		WithContext(ctx).
		Model(&Notification{}).
		Where(&Notification{DocumentID: documentID}).
		Find(&notifications)

	if res.Error != nil {
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	return notifications, nil
}

// Count notifications for a given document
func (db *NotificationDB) Count(ctx context.Context, documentID uuid.UUID) (int64, error) {
	var count int64

	res := db.Conn.
		WithContext(ctx).
		Model(&Notification{}).
		Where(&Notification{DocumentID: documentID}).
		Count(&count)

	if res.Error != nil {
		return 0, status.Error(codes.Internal, res.Error.Error())
	}

	return count, nil
}

// Count all notifications for a given user
func (db *NotificationDB) CountAll(ctx context.Context, userID uuid.UUID) (int64, error) {
	var count int64

	res := db.Conn.
		WithContext(ctx).
		Model(&Notification{}).
		Where(&Notification{UserID: userID}).
		Count(&count)

	if res.Error != nil {
		return 0, status.Error(codes.Internal, res.Error.Error())
	}

	return count, nil
}

func (db *NotificationDB) FindAllForUser(ctx context.Context, userID uuid.UUID) ([]Notification, error) {
	var notifications []Notification

	res := db.Conn.
		WithContext(ctx).
		Model(&Notification{}).
		Where(&Notification{UserID: userID}).
		Find(&notifications)

	if res.Error != nil {
		return nil, status.Error(codes.Internal, res.Error.Error())
	}

	return notifications, nil
}
