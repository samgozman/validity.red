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

type NotificationPayload struct {
	Date time.Time `json:"date" binding:"required"`
}

type NotificationModifyPayload struct {
	ID         string `uri:"id" binding:"required,uuid"`
	DocumentId string `uri:"documentId" binding:"required,uuid"`
}

// Call Create method on Notification in `document-service`
func (app *Config) documentNotificationCreate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uri := struct {
		DocumentId string `uri:"documentId" binding:"required,uuid"`
	}{}

	// get userId from context
	userId, _ := c.Get("UserId")
	// Validate inputs
	if err := c.BindUri(&uri); err != nil {
		c.Error(ErrInvalidInputs)
		return
	}
	notificationPayload := NotificationPayload{}
	if err := c.BindJSON(&notificationPayload); err != nil {
		c.Error(ErrInvalidInputs)
		return
	}

	// call service
	_, err := app.documentsClient.notificationService.Create(ctx, &document.NotificationCreateRequest{
		NotificationEntry: &document.Notification{
			DocumentID: uri.DocumentId,
			Date:       timestamppb.New(notificationPayload.Date),
		},
		UserID: userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::notification::Create method:", err)
		c.Error(err)
		return
	}

	go app.updateIcsCalendar(userId.(string))
	c.Status(http.StatusCreated)
}

// Call Delete method on Notification in `document-service`
func (app *Config) documentNotificationDelete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uri := NotificationModifyPayload{}

	// get userId from context
	userId, _ := c.Get("UserId")
	// Validate inputs
	if err := c.BindUri(&uri); err != nil {
		c.Error(ErrInvalidInputs)
		return
	}

	// call service
	_, err := app.documentsClient.notificationService.Delete(ctx, &document.NotificationCreateRequest{
		NotificationEntry: &document.Notification{
			ID:         uri.ID,
			DocumentID: uri.DocumentId,
		},
		UserID: userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::notification::Delete method:", err)
		c.Error(err)
		return
	}

	go app.updateIcsCalendar(userId.(string))
	c.Status(http.StatusOK)
}

// Call GetAll method on Notification in `document-service`
func (app *Config) documentNotificationGetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	uri := struct {
		DocumentId string `uri:"documentId" binding:"required,uuid"`
	}{}

	// get userId from context
	userId, _ := c.Get("UserId")
	// Validate inputs
	if err := c.BindUri(&uri); err != nil {
		c.Error(ErrInvalidInputs)
		return
	}

	// call service
	res, err := app.documentsClient.notificationService.GetAll(ctx, &document.NotificationsRequest{
		DocumentID: uri.DocumentId,
		UserID:     userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::notification::GetAll method:", err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, struct {
		Notifications []*document.NotificationJSON `json:"notifications"`
	}{
		Notifications: utils.ConvertNotificationsToJSON(res.Notifications),
	})
}
