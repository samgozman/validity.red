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

// Call Create method on Notification in `document-service`
func (app *Config) documentNotificationCreate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")
	documentId := c.Param("documentId")
	notificationPayload := decodeJSON[NotificationPayload](c)

	var payload jsonResponse

	// call service
	res, err := app.documentsClient.notificationService.Create(ctx, &document.NotificationCreateRequest{
		NotificationEntry: &document.Notification{
			DocumentID: documentId,
			Date:       timestamppb.New(notificationPayload.Date),
		},
		UserID: userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::notification::Create method:", err)
		payload.Error = true
		payload.Message = err.Error()
		c.JSON(http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = res.Result

	c.JSON(http.StatusCreated, payload)
}

// Call Edit method on Notification in `document-service`
func (app *Config) documentNotificationEdit(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")
	documentId := c.Param("documentId")
	id := c.Param("id")
	notificationPayload := decodeJSON[NotificationPayload](c)

	var payload jsonResponse

	// call service
	res, err := app.documentsClient.notificationService.Edit(ctx, &document.NotificationCreateRequest{
		NotificationEntry: &document.Notification{
			ID:         id,
			DocumentID: documentId,
			Date:       timestamppb.New(notificationPayload.Date),
		},
		UserID: userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::notification::Edit method:", err)
		payload.Error = true
		payload.Message = err.Error()
		c.JSON(http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = res.Result

	c.JSON(http.StatusCreated, payload)
}

// Call Delete method on Notification in `document-service`
func (app *Config) documentNotificationDelete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")
	documentId := c.Param("documentId")
	id := c.Param("id")

	var payload jsonResponse

	// call service
	res, err := app.documentsClient.notificationService.Delete(ctx, &document.NotificationCreateRequest{
		NotificationEntry: &document.Notification{
			ID:         id,
			DocumentID: documentId,
		},
		UserID: userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::notification::Delete method:", err)
		payload.Error = true
		payload.Message = err.Error()
		c.JSON(http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = res.Result

	c.JSON(http.StatusOK, payload)
}

// Call GetAll method on Notification in `document-service`
func (app *Config) documentNotificationGetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")
	documentId := c.Param("documentId")

	var payload jsonResponse

	// call service
	res, err := app.documentsClient.notificationService.GetAll(ctx, &document.NotificationsRequest{
		DocumentID: documentId,
		UserID:     userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::notification::GetAll method:", err)
		payload.Error = true
		payload.Message = err.Error()
		c.JSON(http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = res.Result
	payload.Data = struct {
		Notifications []*document.NotificationJSON `json:"notifications"`
	}{
		Notifications: utils.ConvertNotficationsToJSON(res.Notifications),
	}

	c.JSON(http.StatusOK, payload)
}
