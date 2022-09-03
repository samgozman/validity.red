package utils

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samgozman/validity.red/document/internal/models/notification"
	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestConvertNotficationsToProtoFormat(t *testing.T) {
	type args struct {
		n *[]notification.Notification
	}

	expectedId, _ := uuid.Parse("884a2112-09a7-4469-bf13-9e9a25b58eab")
	expectedDocumentId, _ := uuid.Parse("7a14e144-120c-4e1e-9447-ece46378c1dd")

	originalNotification := notification.Notification{
		ID:         expectedId,
		DocumentID: expectedDocumentId,
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
			if got := ConvertNotficationsToProtoFormat(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertNotficationsToProtoFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
