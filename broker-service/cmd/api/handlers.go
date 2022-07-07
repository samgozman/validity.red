package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/samgozman/validity.red/broker/proto/logs"
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
	UserID      string    `json:"userId"`
	Type        int32     `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ExpiresAt   time.Time `json:"expiresAt"`
}

type NotificationPayload struct {
	ID         string    `json:"id"`
	DocumentID string    `json:"documentId"`
	UserID     string    `json:"userId"`
	Date       time.Time `json:"date"`
}

// Single point to communicate with services
func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "UserRegister":
		app.userRegister(w, requestPayload.Register)
	case "UserLogin":
		app.userLogin(w, requestPayload.Auth)
	case "DocumentCreate":
		app.documentCreate(w, requestPayload.Document)
	case "DocumentEdit":
		app.documentEdit(w, requestPayload.Document)
	case "DocumentDelete":
		app.documentDelete(w, requestPayload.Document)
	case "DocumentGetOne":
		app.documentGetOne(w, requestPayload.Document)
	case "NotificationCreate":
		app.documentNotificationCreate(w, requestPayload.Notification)
	case "NotificationEdit":
		app.documentNotificationEdit(w, requestPayload.Notification)
	case "NotificationDelete":
		app.documentNotificationDelete(w, requestPayload.Notification)
	default:
		app.errorJSON(w, errors.New("invalid action"))
		go app.logger.LogWarn(&logs.Log{
			Service: "broker-service",
			Message: fmt.Sprintf("Invalid action: %s", requestPayload.Action),
		})
	}
}
