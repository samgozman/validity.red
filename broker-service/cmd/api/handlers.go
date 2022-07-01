package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/samgozman/validity.red/broker/proto/document"
	"github.com/samgozman/validity.red/broker/proto/logs"
	"github.com/samgozman/validity.red/broker/proto/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RequestPayload struct {
	Action   string          `json:"action"`
	Auth     AuthPayload     `json:"auth,omitempty"`
	Register RegisterPayload `json:"register,omitempty"`
	Document DocumentPayload `json:"document,omitempty"`
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
	default:
		app.errorJSON(w, errors.New("invalid action"))
		go app.logger.LogWarn(&logs.Log{
			Service: "broker-service",
			Message: fmt.Sprintf("Invalid action: %s", requestPayload.Action),
		})
	}
}

// Call Register method on `user-service`
func (app *Config) userRegister(w http.ResponseWriter, registerPayload RegisterPayload) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.usersClient.userService.Register(ctx, &user.RegisterRequest{
		RegisterEntry: &user.Register{
			Email:    registerPayload.Email,
			Password: registerPayload.Password,
		},
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "user-service",
			Message: "Error on calling Register method",
			Error:   err.Error(),
		})
		app.errorJSON(w, err)
		return
	}

	// TODO: Send verification email

	var payload jsonResponse
	payload.Error = false
	payload.Message = res.Result

	go app.logger.LogInfo(&logs.Log{
		Service: "user-service",
		Message: res.Result,
	})

	app.writeJSON(w, http.StatusCreated, payload)
}

// Call Login method on `user-service`
func (app *Config) userLogin(w http.ResponseWriter, authPayload AuthPayload) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.usersClient.authService.Login(ctx, &user.AuthRequest{
		AuthEntry: &user.Auth{
			Email:    authPayload.Email,
			Password: authPayload.Password,
		},
	})
	if err != nil {
		go app.logger.LogWarn(&logs.Log{
			Service: "user-service",
			Message: "Error on calling Login method",
			Error:   err.Error(),
		})
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = res.Result

	go app.logger.LogInfo(&logs.Log{
		Service: "user-service",
		Message: res.Result,
	})

	// write jwt token
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   res.Token,
		Expires: time.Unix(res.TokenExpiresAt, 0),
	})

	app.writeJSON(w, http.StatusAccepted, payload)
}

// Call Create method on `document-service`
func (app *Config) documentCreate(w http.ResponseWriter, documentPayload DocumentPayload) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.documentsClient.documentService.Create(ctx, &document.DocumentCreateRequest{
		DocumentEntry: &document.Document{
			// TODO: get user id from jwt token!
			UserID:      documentPayload.UserID,
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
func (app *Config) documentEdit(w http.ResponseWriter, documentPayload DocumentPayload) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := app.documentsClient.documentService.Edit(ctx, &document.DocumentCreateRequest{
		DocumentEntry: &document.Document{
			ID: documentPayload.ID,
			// TODO: get user id from jwt token!
			UserID:      documentPayload.UserID,
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
