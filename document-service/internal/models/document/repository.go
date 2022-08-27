package document

import (
	"context"

	"github.com/google/uuid"
)

type DocumentRepository interface {
	InsertOne(ctx context.Context, d *Document) error
	UpdateOne(ctx context.Context, d *Document) error
	DeleteOne(ctx context.Context, d *Document) error
	FindOne(ctx context.Context, d *Document) error
	Exists(ctx context.Context, d *Document) (bool, error)
	FindAll(ctx context.Context, userId uuid.UUID) ([]Document, error)
}
