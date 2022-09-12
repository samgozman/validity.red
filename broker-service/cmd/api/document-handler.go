package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samgozman/validity.red/broker/internal/utils"
	"github.com/samgozman/validity.red/broker/proto/document"
	"github.com/samgozman/validity.red/broker/proto/logs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DocumentCreate struct {
	Type        int32     `json:"type"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	ExpiresAt   time.Time `json:"expiresAt" binding:"required"`
}

type DocumentEdit struct {
	ID          string    `json:"id" binding:"required"`
	Type        int32     `json:"type"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	ExpiresAt   time.Time `json:"expiresAt" binding:"required"`
}

// Call Create method on `document-service`
func (app *Config) documentCreate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")

	var payload jsonResponse
	documentPayload := decodeJSON[DocumentCreate](c)

	// call service
	res, err := app.documentsClient.documentService.Create(ctx, &document.DocumentCreateRequest{
		DocumentEntry: &document.Document{
			UserID:      userId.(string),
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

		payload.Error = true
		payload.Message = err.Error()
		c.JSON(http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = res.Result
	payload.Data = struct {
		DocumentId string `json:"documentId"`
	}{
		DocumentId: res.DocumentId,
	}

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: res.Result,
	})

	c.JSON(http.StatusCreated, payload)
}

// Call Edit method on `document-service`
func (app *Config) documentEdit(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")

	var payload jsonResponse
	documentPayload := decodeJSON[DocumentEdit](c)

	// call service
	res, err := app.documentsClient.documentService.Edit(ctx, &document.DocumentCreateRequest{
		DocumentEntry: &document.Document{
			ID:          documentPayload.ID,
			UserID:      userId.(string),
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

		payload.Error = true
		payload.Message = err.Error()

		c.JSON(http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = res.Result

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: res.Result,
	})

	c.JSON(http.StatusCreated, payload)
}

// Call Delete method on `document-service`
func (app *Config) documentDelete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")
	// get documentId from param :documentId
	documentId := c.Param("documentId")

	var payload jsonResponse

	// call service
	res, err := app.documentsClient.documentService.Delete(ctx, &document.DocumentRequest{
		DocumentID: documentId,
		UserID:     userId.(string),
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling Delete method",
			Error:   err.Error(),
		})

		payload.Error = true
		payload.Message = err.Error()

		c.JSON(http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = res.Result

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: res.Result,
	})

	c.JSON(http.StatusOK, payload)
}

// Call GetOne method on `document-service`
func (app *Config) documentGetOne(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")
	// get documentId from param :documentId
	documentId := c.Param("documentId")

	var payload jsonResponse

	// call service
	res, err := app.documentsClient.documentService.GetOne(ctx, &document.DocumentRequest{
		DocumentID: documentId,
		UserID:     userId.(string),
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling GetOne method",
			Error:   err.Error(),
		})

		payload.Error = true
		payload.Message = err.Error()

		c.JSON(http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = res.Result
	payload.Data = struct {
		Document *document.DocumentJSON `json:"document"`
	}{
		Document: &document.DocumentJSON{
			ID:          res.Document.ID,
			UserID:      res.Document.UserID,
			Title:       res.Document.Title,
			Type:        res.Document.Type,
			Description: res.Document.Description,
			ExpiresAt:   utils.ParseProtobufDateToString(res.Document.ExpiresAt),
		},
	}

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: res.Result,
	})

	c.JSON(http.StatusOK, payload)
}

// Call GetAll method on `document-service`
func (app *Config) documentGetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")

	var payload jsonResponse

	// call service
	res, err := app.documentsClient.documentService.GetAll(ctx, &document.DocumentsRequest{
		UserID: userId.(string),
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling GetAll method",
			Error:   err.Error(),
		})

		payload.Error = true
		payload.Message = err.Error()

		c.JSON(http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = res.Result
	payload.Data = struct {
		Documents []*document.DocumentJSON `json:"documents"`
	}{
		Documents: utils.ConvertDocumentsToJSON(res.Documents),
	}

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: res.Result,
	})

	c.JSON(http.StatusOK, payload)
}

// TODO: Cache this route
func (app *Config) documentGetStatistics(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")

	var payload jsonResponse
	var statistics struct {
		TotalDocuments     int64                          `json:"totalDocuments"`
		TotalNotifications int64                          `json:"totalNotifications"`
		UsedTypes          []*document.DocumentTypesCount `json:"usedTypes"`
		LatestDocuments    []*document.DocumentJSON       `json:"latestDocuments"`
	}

	// call services
	getStats, err := app.documentsClient.documentService.GetUserStatistics(ctx, &document.DocumentsRequest{
		UserID: userId.(string),
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling GetOne method with GetUserStatistics",
			Error:   err.Error(),
		})

		payload.Error = true
		payload.Message = err.Error()

		c.JSON(http.StatusBadRequest, payload)
		return
	}

	statistics.TotalDocuments = getStats.Total
	statistics.LatestDocuments = utils.ConvertDocumentsToJSON(getStats.LatestDocuments)
	statistics.UsedTypes = getStats.Types

	getIDs, err := app.documentsClient.documentService.GetIDs(ctx, &document.DocumentsRequest{
		UserID: userId.(string),
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "document-service",
			Message: "Error on calling GetOne method with GetIDs",
			Error:   err.Error(),
		})

		payload.Error = true
		payload.Message = err.Error()

		c.JSON(http.StatusBadRequest, payload)
		return
	}

	if len(getIDs.Ids) > 0 {
		totalNotificationsCount, err := app.documentsClient.notificationService.Count(ctx, &document.NotificationsCountRequest{
			UserID:      userId.(string),
			DocumentIDs: getIDs.Ids,
		})
		if err != nil {
			go app.logger.LogWarn(&logs.Log{
				Service: "document-service",
				Message: "Error on calling GetOne method with Count",
				Error:   err.Error(),
			})

			payload.Error = true
			payload.Message = err.Error()

			c.JSON(http.StatusBadRequest, payload)
			return
		}

		statistics.TotalNotifications = totalNotificationsCount.Count
	}

	msg := fmt.Sprintf("User '%s' successfully called documentGetStatistics method", userId.(string))
	payload.Error = false
	payload.Message = msg
	payload.Data = statistics

	go app.logger.LogInfo(&logs.Log{
		Service: "document-service",
		Message: msg,
	})

	c.JSON(http.StatusOK, payload)
}
