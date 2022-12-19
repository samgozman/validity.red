package document

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestDocument_Validate(t *testing.T) {
	tests := []struct {
		name     string
		document Document
		wantErr  bool
	}{
		{
			name:     "user_id null check",
			document: Document{Title: "t", Description: "d", ExpiresAt: time.Now()},
			wantErr:  true,
		},
		{
			name:     "expires_at null check",
			document: Document{UserID: uuid.New(), Title: "t", Description: "d"},
			wantErr:  true,
		},
		{
			name:     "should pass",
			document: Document{UserID: uuid.New(), Title: "t", Description: "d", ExpiresAt: time.Now()},
			wantErr:  false,
		},
		// TODO: add test case for description length
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Document{
				UserID:        tt.document.UserID,
				Type:          tt.document.Type,
				Title:         tt.document.Title,
				Description:   tt.document.Description,
				ExpiresAt:     tt.document.ExpiresAt,
				Notifications: tt.document.Notifications,
				CreatedAt:     tt.document.CreatedAt,
				UpdatedAt:     tt.document.UpdatedAt,
			}
			if err := d.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Document.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDocument_BeforeCreate(t *testing.T) {
	type args struct {
		tx *gorm.DB
	}

	EncryptionKey = []byte("8dHWTNSAsGaaD7JbqVubF1aWVWGJYF3q")

	tests := []struct {
		name     string
		document Document
		args     args
		wantErr  bool
	}{
		{
			name:     "should pass and create new id",
			document: Document{UserID: uuid.New(), Title: "t", Description: "d", ExpiresAt: time.Now()},
			args:     args{tx: &gorm.DB{}},
			wantErr:  false,
		},
		{
			name:     "should fail on user_id validation",
			document: Document{Title: "t", Description: "d", ExpiresAt: time.Now()},
			args:     args{tx: &gorm.DB{}},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.document.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("Document.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.document.ID == uuid.Nil {
				t.Errorf("Document.BeforeCreate() want ID not nil")
			}
		})
	}
}

// TODO: test methods with DB call
