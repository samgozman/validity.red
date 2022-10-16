package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samgozman/validity.red/broker/internal/utils"
	"github.com/samgozman/validity.red/broker/proto/document"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ? Maybe use alpha-num-unicode rule for string fields?

type DocumentCreate struct {
	Type        int32     `json:"type" binding:"required,number"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	ExpiresAt   time.Time `json:"expiresAt" binding:"required"`
}

type DocumentEdit struct {
	ID          string    `json:"id" binding:"required,uuid"`
	Type        int32     `json:"type" binding:"required,number"`
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
	documentPayload := DocumentCreate{}
	if err := c.BindJSON(&documentPayload); err != nil {
		payload.Error = true
		payload.Message = "Invalid document payload."
		c.AbortWithStatusJSON(http.StatusBadRequest, payload)
		return
	}

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
		log.Println("Error on calling document-service::Create method:", err)

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

	c.JSON(http.StatusCreated, payload)
}

// Call Edit method on `document-service`
func (app *Config) documentEdit(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")

	var payload jsonResponse
	documentPayload := DocumentEdit{}
	if err := c.BindJSON(&documentPayload); err != nil {
		payload.Error = true
		payload.Message = "Invalid document payload."
		c.AbortWithStatusJSON(http.StatusBadRequest, payload)
		return
	}

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
		log.Println("Error on calling document-service::Edit method:", err)

		payload.Error = true
		payload.Message = err.Error()

		c.JSON(http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = res.Result

	go app.updateIcsCalendar(userId.(string))
	c.JSON(http.StatusCreated, payload)
}

// Call Delete method on `document-service`
func (app *Config) documentDelete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var payload jsonResponse

	// get userId from context
	userId, _ := c.Get("UserId")
	uri := struct {
		DocumentId string `uri:"documentId" binding:"required,uuid"`
	}{}
	if err := c.BindUri(&uri); err != nil {
		payload.Error = true
		payload.Message = "Invalid documentId."
		c.AbortWithStatusJSON(http.StatusBadRequest, payload)
	}

	// call service
	res, err := app.documentsClient.documentService.Delete(ctx, &document.DocumentRequest{
		DocumentID: uri.DocumentId,
		UserID:     userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::Delete method:", err)

		payload.Error = true
		payload.Message = err.Error()

		c.JSON(http.StatusBadRequest, payload)
		return
	}

	payload.Error = false
	payload.Message = res.Result

	go app.updateIcsCalendar(userId.(string))
	c.JSON(http.StatusOK, payload)
}

// Call GetOne method on `document-service`
func (app *Config) documentGetOne(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var payload jsonResponse

	// get userId from context
	userId, _ := c.Get("UserId")
	uri := struct {
		DocumentId string `uri:"documentId" binding:"required,uuid"`
	}{}
	if err := c.BindUri(&uri); err != nil {
		payload.Error = true
		payload.Message = "Invalid documentId."
		c.AbortWithStatusJSON(http.StatusBadRequest, payload)
	}

	// call service
	res, err := app.documentsClient.documentService.GetOne(ctx, &document.DocumentRequest{
		DocumentID: uri.DocumentId,
		UserID:     userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::GetOne method:", err)

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
		log.Println("Error on calling document-service::GetAll method:", err)

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
		log.Println("Error on calling document-service::GetUserStatistics method:", err)

		payload.Error = true
		payload.Message = err.Error()

		c.JSON(http.StatusBadRequest, payload)
		return
	}

	statistics.TotalDocuments = getStats.Total
	statistics.LatestDocuments = utils.ConvertDocumentsToJSON(getStats.LatestDocuments)
	statistics.UsedTypes = getStats.Types

	totalNotificationsCount, err := app.documentsClient.notificationService.CountAll(
		ctx,
		&document.NotificationsAllRequest{
			UserID: userId.(string),
		},
	)
	if err != nil {
		log.Println("Error on calling document-service::notification::CountAll method:", err)

		payload.Error = true
		payload.Message = err.Error()

		c.JSON(http.StatusBadRequest, payload)
		return
	}
	statistics.TotalNotifications = totalNotificationsCount.Count

	msg := fmt.Sprintf("User '%s' successfully called documentGetStatistics method", userId.(string))
	payload.Error = false
	payload.Message = msg
	payload.Data = statistics

	c.JSON(http.StatusOK, payload)
}
