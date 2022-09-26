package utils

import (
	"time"

	"github.com/samgozman/validity.red/broker/proto/calendar"
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

func ConvertDocumentsToJSON(ds []*document.Document) []*document.DocumentJSON {
	var djs []*document.DocumentJSON
	for _, d := range ds {
		djs = append(djs, &document.DocumentJSON{
			ID:          d.ID,
			UserID:      d.UserID,
			Title:       d.Title,
			Type:        d.Type,
			Description: d.Description,
			ExpiresAt:   ParseProtobufDateToString(d.ExpiresAt),
		})
	}

	return djs
}

func ConvertCalendarToJSON(cl []*calendar.CalendarEntity) []*calendar.CalendarEntityJSON {
	var cldr []*calendar.CalendarEntityJSON
	for _, n := range cl {
		cldr = append(cldr, &calendar.CalendarEntityJSON{
			DocumentID:       n.DocumentID,
			DocumentTitle:    n.DocumentTitle,
			NotificationDate: ParseProtobufDateToString(n.NotificationDate),
			ExpiresAt:        ParseProtobufDateToString(n.ExpiresAt),
		})
	}

	return cldr
}
