package documentmocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/document"
	proto "github.com/samgozman/validity.red/document/proto"
	"gorm.io/gorm"
)

type DocumentDBTest struct {
	Conn *gorm.DB
}

func NewDocumentDBTest(db *gorm.DB) *DocumentDBTest {
	return &DocumentDBTest{
		Conn: db,
	}
}

func (db *DocumentDBTest) InsertOne(ctx context.Context, d *document.Document) error {
	return nil
}

func (db *DocumentDBTest) UpdateOne(ctx context.Context, d *document.Document) error {
	return nil
}

func (db *DocumentDBTest) DeleteOne(ctx context.Context, d *document.Document) error {
	return nil
}

func (db *DocumentDBTest) FindOne(ctx context.Context, d *document.Document) error {
	d.Type = proto.Type_DEFAULT_DOCUMENT.Enum()
	return nil
}

func (db *DocumentDBTest) Exists(ctx context.Context, d *document.Document) (bool, error) {
	if d.ID.String() != "434377cf-7509-4cc0-9895-0afa683f0e56" {
		return false, nil
	}

	return true, nil
}

func (db *DocumentDBTest) FindAll(ctx context.Context, userID uuid.UUID) ([]document.Document, error) {
	var documents []document.Document

	return documents, nil
}

func (db *DocumentDBTest) Count(ctx context.Context, userID uuid.UUID) (int64, error) {
	return 0, nil
}

func (db *DocumentDBTest) CountTypes(ctx context.Context, userID uuid.UUID) ([]*proto.DocumentTypesCount, error) {
	var types []*proto.DocumentTypesCount
	return types, nil
}

func (db *DocumentDBTest) FindLatest(ctx context.Context, userID uuid.UUID, limit int) ([]document.Document, error) {
	var documents []document.Document
	return documents, nil
}
