package utils

import (
	"github.com/samgozman/validity.red/document/internal/models/notification"
	proto "github.com/samgozman/validity.red/document/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertNotficationsToProtoFormat(n *[]notification.Notification) []*proto.Notification {
	var result []*proto.Notification

	for _, n := range *n {
		result = append(result, &proto.Notification{
			ID:         n.ID.String(),
			DocumentID: n.DocumentID.String(),
			Date:       timestamppb.New(n.Date),
		})
	}

	return result
}
