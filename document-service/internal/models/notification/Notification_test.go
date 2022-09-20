package notification

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestNotification_Prepare(t *testing.T) {
	tests := []struct {
		name   string
		fields *Notification
	}{
		{
			name:   "should pass",
			fields: &Notification{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notification{
				ID:         tt.fields.ID,
				DocumentID: tt.fields.DocumentID,
				Date:       tt.fields.Date,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			n.Prepare()
			if n.CreatedAt.IsZero() || n.UpdatedAt.IsZero() {
				t.Errorf("Notification.Prepare() want CreatedAt and UpdatedAt not zero")
			}
		})
	}
}

func TestNotification_Validate(t *testing.T) {
	tests := []struct {
		name    string
		fields  *Notification
		wantErr bool
	}{
		{
			name:    "fail if userID is empty",
			fields:  &Notification{UserID: uuid.UUID{}, DocumentID: uuid.New(), Date: time.Now()},
			wantErr: true,
		},
		{
			name:    "fail if documentID is empty",
			fields:  &Notification{UserID: uuid.New(), DocumentID: uuid.UUID{}, Date: time.Now()},
			wantErr: true,
		},
		{
			name:    "fail if Date is empty",
			fields:  &Notification{UserID: uuid.New(), DocumentID: uuid.New()},
			wantErr: true,
		},
		{
			name:    "should pass",
			fields:  &Notification{UserID: uuid.New(), DocumentID: uuid.New(), Date: time.Now()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notification{
				ID:         tt.fields.ID,
				UserID:     tt.fields.UserID,
				DocumentID: tt.fields.DocumentID,
				Date:       tt.fields.Date,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			if err := n.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Notification.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotification_BeforeCreate(t *testing.T) {
	type args struct {
		tx *gorm.DB
	}

	a := args{
		tx: &gorm.DB{},
	}

	tests := []struct {
		name    string
		fields  *Notification
		args    args
		wantErr bool
	}{
		{
			name:    "should fail on validation",
			fields:  &Notification{UserID: uuid.New(), DocumentID: uuid.UUID{}, Date: time.Now()},
			args:    a,
			wantErr: true,
		},
		{
			name:    "should pass",
			fields:  &Notification{UserID: uuid.New(), DocumentID: uuid.New(), Date: time.Now()},
			args:    a,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notification{
				ID:         tt.fields.ID,
				UserID:     tt.fields.UserID,
				DocumentID: tt.fields.DocumentID,
				Date:       tt.fields.Date,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			if err := n.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("Notification.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if n.ID.String() == "" {
				t.Error("Notification.ID is empty")
			}
		})
	}
}
