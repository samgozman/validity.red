package notification_mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/notification"
	"gorm.io/gorm"
)

type NotificationDBTest struct {
	Conn *gorm.DB
}

func NewNotificationDBTest(db *gorm.DB) *NotificationDBTest {
	return &NotificationDBTest{
		Conn: db,
	}
}

func (db *NotificationDBTest) InsertOne(ctx context.Context, n *notification.Notification) error {
	return nil
}

func (db *NotificationDBTest) UpdateOne(ctx context.Context, n *notification.Notification) error {
	return nil
}

func (db *NotificationDBTest) DeleteOne(ctx context.Context, n *notification.Notification) error {
	return nil
}

func (db *NotificationDBTest) FindAll(ctx context.Context, documentID uuid.UUID) ([]notification.Notification, error) {
	var notifications []notification.Notification
	return notifications, nil
}

func (db *NotificationDBTest) Count(ctx context.Context, documentID uuid.UUID) (int64, error) {
	return 0, nil
}
