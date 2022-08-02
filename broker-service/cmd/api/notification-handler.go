package main

import (
	"context"
	"net/http"
	"time"

	"github.com/samgozman/validity.red/broker/proto/document"
	"github.com/samgozman/validity.red/broker/proto/logs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Call Create method on Notification in `document-service`
func (app *Config) documentNotificationCreate(
	w http.ResponseWriter,
	notificationPayload NotificationPayload,
	userId string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.documentsClient.notificationService.Create(ctx, &document.NotificationCreateRequest{
		NotificationEntry: &document.Notification{
			DocumentID: notificationPayload.DocumentID,
			Date:       timestamppb.New(notificationPayload.Date),
		},
		UserID: userId,
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling Notification.Create method",
			Error:   err.Error(),
		})
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = res.Result

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: res.Result,
	})

	app.writeJSON(w, http.StatusCreated, payload)
}

// Call Edit method on Notification in `document-service`
func (app *Config) documentNotificationEdit(
	w http.ResponseWriter,
	notificationPayload NotificationPayload,
	userId string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.documentsClient.notificationService.Edit(ctx, &document.NotificationCreateRequest{
		NotificationEntry: &document.Notification{
			ID:         notificationPayload.ID,
			DocumentID: notificationPayload.DocumentID,
			Date:       timestamppb.New(notificationPayload.Date),
		},
		UserID: userId,
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling Notification.Edit method",
			Error:   err.Error(),
		})
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = res.Result

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: res.Result,
	})

	app.writeJSON(w, http.StatusCreated, payload)
}

// Call Delete method on Notification in `document-service`
func (app *Config) documentNotificationDelete(
	w http.ResponseWriter,
	notificationPayload NotificationPayload,
	userId string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.documentsClient.notificationService.Delete(ctx, &document.NotificationCreateRequest{
		NotificationEntry: &document.Notification{
			ID:         notificationPayload.ID,
			DocumentID: notificationPayload.DocumentID,
		},
		UserID: userId,
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling Notification.Delete method",
			Error:   err.Error(),
		})
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = res.Result

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: res.Result,
	})

	app.writeJSON(w, http.StatusCreated, payload)
}

// Call GetAll method on Notification in `document-service`
func (app *Config) documentNotificationGetAll(
	w http.ResponseWriter,
	notificationPayload NotificationPayload,
	userId string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.documentsClient.notificationService.GetAll(ctx, &document.NotificationsRequest{
		DocumentID: notificationPayload.DocumentID,
		UserID:     userId,
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling Notification.GetAll method",
			Error:   err.Error(),
		})
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = res.Result
	payload.Data = struct {
		Notifications []*document.Notification `json:"notifications"`
	}{
		Notifications: res.Notifications,
	}

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: res.Result,
	})

	app.writeJSON(w, http.StatusCreated, payload)
}
