package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestPayload struct {
	Action       string              `json:"action"`
	Auth         AuthPayload         `json:"auth,omitempty"`
	Register     RegisterPayload     `json:"register,omitempty"`
	Document     DocumentPayload     `json:"document,omitempty"`
	Notification NotificationPayload `json:"notification,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DocumentPayload struct {
	ID          string    `json:"id"`
	Type        int32     `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

type NotificationPayload struct {
	ID         string    `json:"id"`
	DocumentID string    `json:"documentId"`
	Date       time.Time `json:"date"`
}

var (
	ErrInvalidAction = errors.New("invalid action")
)

// Single point to communicate with services
func (app *Config) HandleSubmission(c *gin.Context) {
	requestPayload := decodeJSON[RequestPayload](c)

	switch requestPayload.Action {
	case "UserRefreshToken":
		app.userRefreshToken(c)
	case "DocumentCreate":
		app.documentCreate(c, requestPayload.Document)
	case "DocumentEdit":
		app.documentEdit(c, requestPayload.Document)
	case "DocumentDelete":
		app.documentDelete(c, requestPayload.Document)
	case "DocumentGetOne":
		app.documentGetOne(c, requestPayload.Document)
	case "DocumentGetAll":
		app.documentGetAll(c)
	case "NotificationCreate":
		app.documentNotificationCreate(c, requestPayload.Notification)
	case "NotificationEdit":
		app.documentNotificationEdit(c, requestPayload.Notification)
	case "NotificationDelete":
		app.documentNotificationDelete(c, requestPayload.Notification)
	case "NotificationGetAll":
		app.documentNotificationGetAll(c, requestPayload.Notification)
	default:
		c.AbortWithError(http.StatusBadRequest, ErrInvalidAction)
	}
}
