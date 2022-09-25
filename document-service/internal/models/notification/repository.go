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
	Count(ctx context.Context, documentID uuid.UUID) (int64, error)
	CountAll(ctx context.Context, userID uuid.UUID) (int64, error)
	FindAllForUser(ctx context.Context, params NotificationFindAllForUser) ([]Notification, error)
}
