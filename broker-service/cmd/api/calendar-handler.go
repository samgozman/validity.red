package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samgozman/validity.red/broker/internal/utils"
	"github.com/samgozman/validity.red/broker/proto/calendar"
	"github.com/samgozman/validity.red/broker/proto/document"
)

// TODO: add pagination by month
func (app *Config) getCalendar(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")

	var payload jsonResponse

	documents, err := app.documentsClient.documentService.GetAll(ctx, &document.DocumentsRequest{
		UserID: userId.(string),
	})
	if err != nil {
		log.Println("Error on calling GetAll method for getCalendar:", err)
		payload.Error = true
		payload.Message = err.Error()
		c.JSON(http.StatusBadRequest, payload)
		return
	}

	notifications, err := app.documentsClient.notificationService.GetAllForUser(ctx, &document.NotificationsAllRequest{
		UserID: userId.(string),
	})
	if err != nil {
		log.Println("Error on calling Notification.GetAllForUser method for getCalendar:", err)
		payload.Error = true
		payload.Message = err.Error()
		c.JSON(http.StatusBadRequest, payload)
		return
	}

	calendarArr := createCalendar(documents.Documents, notifications.Notifications)

	payload.Data = struct {
		Calendar []*calendar.CalendarEntityJSON `json:"calendar"`
	}{
		Calendar: utils.ConvertCalendarToJSON(calendarArr),
	}

	c.JSON(http.StatusOK, payload)
}

// Combine array of documents with array of notifications
// into array of CalendarEntity
func createCalendar(
	documents []*document.Document,
	notifications []*document.Notification,
) []*calendar.CalendarEntity {
	var calendarArr []*calendar.CalendarEntity

	for _, notification := range notifications {
		d := findDocumentByID(documents, notification.DocumentID)
		calendarArr = append(calendarArr, &calendar.CalendarEntity{
			DocumentID:       d.ID,
			NotificationID:   notification.ID,
			DocumentTitle:    d.Title,
			NotificationDate: notification.Date,
			ExpiresAt:        d.ExpiresAt,
		})
	}

	return calendarArr
}

// Find document by ID in array of documents
func findDocumentByID(documents []*document.Document, id string) *document.Document {
	for _, document := range documents {
		if document.ID == id {
			return document
		}
	}

	return &document.Document{}
}
