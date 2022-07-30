package main

import (
	"context"
	"net/http"
	"time"

	"github.com/samgozman/validity.red/broker/proto/document"
	"github.com/samgozman/validity.red/broker/proto/logs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Call Create method on `document-service`
func (app *Config) documentCreate(
	w http.ResponseWriter,
	documentPayload DocumentPayload,
	userId string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.documentsClient.documentService.Create(ctx, &document.DocumentCreateRequest{
		DocumentEntry: &document.Document{
			UserID:      userId,
			Title:       documentPayload.Title,
			Type:        document.Type(documentPayload.Type),
			Description: documentPayload.Description,
			ExpiresAt:   timestamppb.New(documentPayload.ExpiresAt),
		},
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling Create method",
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

// Call Edit method on `document-service`
func (app *Config) documentEdit(
	w http.ResponseWriter,
	documentPayload DocumentPayload,
	userId string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.documentsClient.documentService.Edit(ctx, &document.DocumentCreateRequest{
		DocumentEntry: &document.Document{
			ID:          documentPayload.ID,
			UserID:      userId,
			Title:       documentPayload.Title,
			Type:        document.Type(documentPayload.Type),
			Description: documentPayload.Description,
			ExpiresAt:   timestamppb.New(documentPayload.ExpiresAt),
		},
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling Edit method",
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

// Call Delete method on `document-service`
func (app *Config) documentDelete(
	w http.ResponseWriter,
	documentPayload DocumentPayload,
	userId string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.documentsClient.documentService.Delete(ctx, &document.DocumentRequest{
		DocumentID: documentPayload.ID,
		UserID:     userId,
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling Delete method",
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

	app.writeJSON(w, http.StatusOK, payload)
}

// Call GetOne method on `document-service`
func (app *Config) documentGetOne(
	w http.ResponseWriter,
	documentPayload DocumentPayload,
	userId string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.documentsClient.documentService.GetOne(ctx, &document.DocumentRequest{
		DocumentID: documentPayload.ID,
		UserID:     userId,
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling GetOne method",
			Error:   err.Error(),
		})
		app.errorJSON(w, err)
		return
	}

	// TODO: Convert ExpiresAt to time.Time
	// TODO: Convert Notification.Data to time.Time

	var payload jsonResponse
	payload.Error = false
	payload.Message = res.Result
	payload.Data = struct {
		Document      *document.Document       `json:"documents"`
		Notifications []*document.Notification `json:"notifications"`
	}{
		Document:      res.Document,
		Notifications: res.Notifications,
	}

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: res.Result,
	})

	app.writeJSON(w, http.StatusOK, payload)
}

// Call GetAll method on `document-service`
func (app *Config) documentGetAll(
	w http.ResponseWriter,
	userId string,
) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.documentsClient.documentService.GetAll(ctx, &document.DocumentsRequest{
		UserID: userId,
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling GetOne method",
			Error:   err.Error(),
		})
		app.errorJSON(w, err)
		return
	}

	// TODO: Convert ExpiresAt to time.Time
	// TODO: Convert Notification.Data to time.Time

	var payload jsonResponse
	payload.Error = false
	payload.Message = res.Result
	payload.Data = struct {
		Documents []*document.Document `json:"documents"`
	}{
		Documents: res.Documents,
	}

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: res.Result,
	})

	app.writeJSON(w, http.StatusOK, payload)
}
