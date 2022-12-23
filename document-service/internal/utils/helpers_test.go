package utils

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/document"
	"github.com/samgozman/validity.red/document/internal/models/notification"
	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestConvertNotificationsToProtoFormat(t *testing.T) {
	type args struct {
		n *[]notification.Notification
	}

	expectedID, _ := uuid.Parse("884a2112-09a7-4469-bf13-9e9a25b58eab")
	expectedDocumentID, _ := uuid.Parse("7a14e144-120c-4e1e-9447-ece46378c1dd")

	originalNotification := notification.Notification{
		ID:         expectedID,
		DocumentID: expectedDocumentID,
		Date:       time.Unix(1662204865, 0),
	}

	protoNotification := proto.Notification{
		ID:         "884a2112-09a7-4469-bf13-9e9a25b58eab",
		DocumentID: "7a14e144-120c-4e1e-9447-ece46378c1dd",
		Date: &timestamppb.Timestamp{
			Seconds: 1662204865,
			Nanos:   0,
		},
	}

	tests := []struct {
		name string
		args args
		want []*proto.Notification
	}{
		{
			name: "should convert notifications to proto format",
			args: args{
				n: &[]notification.Notification{
					originalNotification,
				},
			},
			want: []*proto.Notification{
				&protoNotification,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNotificationsToProtoFormat(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNotificationsToProtoFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertDocumentsToProtoFormat(t *testing.T) {
	type args struct {
		d *[]document.Document
	}

	expectedUserID, _ := uuid.Parse("884a2112-09a7-4469-bf13-9e9a25b58eab")
	expectedDocumentID, _ := uuid.Parse("7a14e144-120c-4e1e-9447-ece46378c1dd")

	var docType proto.Type = 1

	originalDocument := document.Document{
		ID:        expectedDocumentID,
		UserID:    expectedUserID,
		Type:      &docType,
		Title:     "Some title",
		ExpiresAt: time.Unix(1662204865, 0),
	}

	expectedProtoDocument := proto.Document{
		ID:     "7a14e144-120c-4e1e-9447-ece46378c1dd",
		UserID: "884a2112-09a7-4469-bf13-9e9a25b58eab",
		Type:   1,
		Title:  "Some title",
		ExpiresAt: &timestamppb.Timestamp{
			Seconds: 1662204865,
			Nanos:   0,
		},
	}

	tests := []struct {
		name string
		args args
		want []*proto.Document
	}{
		{
			name: "should convert notifications to proto format",
			args: args{
				d: &[]document.Document{
					originalDocument,
				},
			},
			want: []*proto.Document{
				&expectedProtoDocument,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertDocumentsToProtoFormat(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertDocumentsToProtoFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
