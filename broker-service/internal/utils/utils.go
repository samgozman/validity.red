package utils

import (
	"time"

	"github.com/samgozman/validity.red/broker/proto/document"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ParseProtobufDateToString(t *timestamppb.Timestamp) string {
	return t.AsTime().Format(time.RFC3339)
}

func ConvertNotficationsToJSON(ns []*document.Notification) []*document.NotificationJSON {
	var njs []*document.NotificationJSON
	for _, n := range ns {
		njs = append(njs, &document.NotificationJSON{
			ID:         n.ID,
			DocumentID: n.DocumentID,
			Date:       ParseProtobufDateToString(n.Date),
		})
	}

	return njs
}
