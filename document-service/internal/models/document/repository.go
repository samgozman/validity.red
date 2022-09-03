package document

import (
	"context"

	"github.com/google/uuid"
	proto "github.com/samgozman/validity.red/document/proto"
)

type DocumentRepository interface {
	InsertOne(ctx context.Context, d *Document) error
	UpdateOne(ctx context.Context, d *Document) error
	DeleteOne(ctx context.Context, d *Document) error
	FindOne(ctx context.Context, d *Document) error
	Exists(ctx context.Context, d *Document) (bool, error)
	FindAll(ctx context.Context, userId uuid.UUID) ([]Document, error)
	Count(ctx context.Context, userId uuid.UUID) (int64, error)
	CountTypes(ctx context.Context, userId uuid.UUID) ([]*proto.DocumentTypesCount, error)
}
