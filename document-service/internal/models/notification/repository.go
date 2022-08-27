package notification

import (
	"context"

	"github.com/google/uuid"
)

type NotificationRepository interface {
	InsertOne(ctx context.Context, n *Notification) error
	UpdateOne(ctx context.Context, n *Notification) error
	DeleteOne(ctx context.Context, n *Notification) error
	FindAll(ctx context.Context, documentID uuid.UUID) ([]Notification, error)
}
