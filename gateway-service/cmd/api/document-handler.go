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

	documentPayload := DocumentCreate{}
	if err := c.BindJSON(&documentPayload); err != nil {
		c.Error(ErrInvalidInputs)
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
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, struct {
		DocumentId string `json:"documentId"`
	}{
		DocumentId: res.DocumentId,
	})
}

// Call Edit method on `document-service`
func (app *Config) documentEdit(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")

	documentPayload := DocumentEdit{}
	if err := c.BindJSON(&documentPayload); err != nil {
		c.Error(ErrInvalidInputs)
		return
	}

	// call service
	_, err := app.documentsClient.documentService.Edit(ctx, &document.DocumentCreateRequest{
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
		c.Error(err)
		return
	}

	go app.updateIcsCalendar(userId.(string))
	c.Status(http.StatusCreated)
}

// Call Delete method on `document-service`
func (app *Config) documentDelete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")
	uri := struct {
		DocumentId string `uri:"documentId" binding:"required,uuid"`
	}{}
	if err := c.BindUri(&uri); err != nil {
		c.Error(ErrInvalidInputs)
		return
	}

	// call service
	_, err := app.documentsClient.documentService.Delete(ctx, &document.DocumentRequest{
		DocumentID: uri.DocumentId,
		UserID:     userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::Delete method:", err)
		c.Error(err)
		return
	}

	go app.updateIcsCalendar(userId.(string))
	c.Status(http.StatusOK)
}

// Call GetOne method on `document-service`
func (app *Config) documentGetOne(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")
	uri := struct {
		DocumentId string `uri:"documentId" binding:"required,uuid"`
	}{}
	if err := c.BindUri(&uri); err != nil {
		c.Error(ErrInvalidInputs)
		return
	}

	// call service
	res, err := app.documentsClient.documentService.GetOne(ctx, &document.DocumentRequest{
		DocumentID: uri.DocumentId,
		UserID:     userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::GetOne method:", err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, struct {
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
	})
}

// Call GetAll method on `document-service`
func (app *Config) documentGetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")

	// call service
	res, err := app.documentsClient.documentService.GetAll(ctx, &document.DocumentsRequest{
		UserID: userId.(string),
	})
	if err != nil {
		log.Println("Error on calling document-service::GetAll method:", err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, struct {
		Documents []*document.DocumentJSON `json:"documents"`
	}{
		Documents: utils.ConvertDocumentsToJSON(res.Documents),
	})
}

// TODO: Cache this route
func (app *Config) documentGetStatistics(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// get userId from context
	userId, _ := c.Get("UserId")

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
		c.Error(err)
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
		c.Error(err)
		return
	}
	statistics.TotalNotifications = totalNotificationsCount.Count

	c.JSON(http.StatusOK, statistics)
}
