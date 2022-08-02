package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/exp/slices"

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
	// Handlers that doesn't require authentication
	WhiteListedHandlers = []string{"UserRegister", "UserLogin"}
)

// Single point to communicate with services
func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var (
		userId string
		token  *http.Cookie
	)
	// Check if requested handler requires authentication
	if !slices.Contains(WhiteListedHandlers, requestPayload.Action) {
		// Get token from cookie
		token, err = r.Cookie("token")
		if err != nil {
			app.errorJSON(w, errors.New("error occured while reading token cookie"), http.StatusUnauthorized)
			return
		}

		// Verify token and decode UserId from it
		userId, err = app.token.Verify(token.Value)
		if err != nil {
			app.errorJSON(w, err, http.StatusUnauthorized)
			return
		}
	}

	switch requestPayload.Action {
	case "UserRegister":
		app.userRegister(w, requestPayload.Register)
	case "UserLogin":
		app.userLogin(w, requestPayload.Auth)
	case "UserRefreshToken":
		app.userRefreshToken(w, userId, token.Value)
	case "DocumentCreate":
		app.documentCreate(w, requestPayload.Document, userId)
	case "DocumentEdit":
		app.documentEdit(w, requestPayload.Document, userId)
	case "DocumentDelete":
		app.documentDelete(w, requestPayload.Document, userId)
	case "DocumentGetOne":
		app.documentGetOne(w, requestPayload.Document, userId)
	case "DocumentGetAll":
		app.documentGetAll(w, userId)
	case "NotificationCreate":
		app.documentNotificationCreate(w, requestPayload.Notification, userId)
	case "NotificationEdit":
		app.documentNotificationEdit(w, requestPayload.Notification, userId)
	case "NotificationDelete":
		app.documentNotificationDelete(w, requestPayload.Notification, userId)
	case "NotificationGetAll":
		app.documentNotificationGetAll(w, requestPayload.Notification, userId)
	default:
		app.errorJSON(w, errors.New("invalid action"))
		go app.logger.LogWarn(&logs.Log{
			Service: "broker-service",
			Message: fmt.Sprintf("Invalid action: %s", requestPayload.Action),
		})
	}
}
