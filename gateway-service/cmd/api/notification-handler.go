package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samgozman/validity.red/broker/internal/utils"
	"github.com/samgozman/validity.red/broker/proto/document"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type notificationPayload struct {
	Date time.Time `json:"date" binding:"required"`
}

type notificationModifyPayload struct {
	ID         string `uri:"id" binding:"required,uuid"`
	DocumentID string `uri:"documentId" binding:"required,uuid"`
}

// Call Create method on Notification in `document-service`.
func (app *Config) documentNotificationCreate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uri := struct {
		DocumentID string `uri:"documentId" binding:"required,uuid"`
	}{}

	// get userID from context
	userID, _ := c.Get("UserId")
	// Validate inputs
	if err := c.BindUri(&uri); err != nil {
		_ = c.Error(ErrInvalidInputs)
		return
	}

	payload := notificationPayload{}
	if err := c.BindJSON(&payload); err != nil {
		_ = c.Error(ErrInvalidInputs)
		return
	}

	// call service
	_, err := app.documentsClient.notificationService.Create(ctx, &document.NotificationCreateRequest{
		NotificationEntry: &document.Notification{
			DocumentID: uri.DocumentID,
			Date:       timestamppb.New(payload.Date),
		},
		UserID: userID.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::notification::Create method:", err)
		_ = c.Error(err)

		return
	}

	go app.updateIcsCalendar(userID.(string))
	c.Status(http.StatusCreated)
}

// Call Delete method on Notification in `document-service`.
func (app *Config) documentNotificationDelete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uri := notificationModifyPayload{}

	// get userID from context
	userID, _ := c.Get("UserId")

	// Validate inputs
	if err := c.BindUri(&uri); err != nil {
		_ = c.Error(ErrInvalidInputs)
		return
	}

	// call service
	_, err := app.documentsClient.notificationService.Delete(ctx, &document.NotificationCreateRequest{
		NotificationEntry: &document.Notification{
			ID:         uri.ID,
			DocumentID: uri.DocumentID,
		},
		UserID: userID.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::notification::Delete method:", err)
		_ = c.Error(err)

		return
	}

	go app.updateIcsCalendar(userID.(string))
	c.Status(http.StatusOK)
}

// Call GetAll method on Notification in `document-service`.
func (app *Config) documentNotificationGetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uri := struct {
		DocumentID string `uri:"documentId" binding:"required,uuid"`
	}{}

	// get userID from context
	userID, _ := c.Get("UserId")

	// Validate inputs
	if err := c.BindUri(&uri); err != nil {
		_ = c.Error(ErrInvalidInputs)
		return
	}

	// call service
	res, err := app.documentsClient.notificationService.GetAll(ctx, &document.NotificationsRequest{
		DocumentID: uri.DocumentID,
		UserID:     userID.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::notification::GetAll method:", err)
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, struct {
		Notifications []*document.NotificationJSON `json:"notifications"`
	}{
		Notifications: utils.ConvertNotificationsToJSON(res.Notifications),
	})
}
