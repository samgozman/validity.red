package utils

import (
	"github.com/samgozman/validity.red/document/internal/models/document"
	"github.com/samgozman/validity.red/document/internal/models/notification"
	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertNotificationsToProtoFormat(n *[]notification.Notification) []*proto.Notification {
	var result = []*proto.Notification{}

	for _, n := range *n {
		result = append(result, &proto.Notification{
			ID:         n.ID.String(),
			DocumentID: n.DocumentID.String(),
			Date:       timestamppb.New(n.Date),
		})
	}

	return result
}

func ConvertDocumentsToProtoFormat(d *[]document.Document) []*proto.Document {
	var result = []*proto.Document{}

	for _, d := range *d {
		result = append(result, &proto.Document{
			ID:          d.ID.String(),
			UserID:      d.UserID.String(),
			Title:       d.Title,
			Type:        d.Type,
			Description: d.Description,
			ExpiresAt:   timestamppb.New(d.ExpiresAt),
		})
	}

	return result
}
