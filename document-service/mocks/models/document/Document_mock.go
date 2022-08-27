package document_mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/document"
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
	return nil
}

// Checks if document is already exists in database
func (db *DocumentDBTest) Exists(ctx context.Context, d *document.Document) (bool, error) {
	return true, nil
}

// Find all documents by UserID
func (db *DocumentDBTest) FindAll(ctx context.Context, userId uuid.UUID) ([]document.Document, error) {
	var documents []document.Document

	return documents, nil
}
