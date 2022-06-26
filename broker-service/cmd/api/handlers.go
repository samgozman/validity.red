package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/samgozman/validity.red/broker/proto/logs"
	"github.com/samgozman/validity.red/broker/proto/user"
)

type RequestPayload struct {
	Action   string          `json:"action"`
	Auth     AuthPayload     `json:"auth,omitempty"`
	Register RegisterPayload `json:"register,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
	// connect to gRPC
	conn, err := connectToUserService()
	if err != nil {
		go app.logger.LogFatal(&logs.Log{
			Service: "user-service",
			Message: "Error on connecting to the user-service",
			Error:   err.Error(),
		})
		app.errorJSON(w, errors.New("service is unavailable. Please try again later"))
		return
	}
	defer conn.Close()

	// create client
	client := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := client.Register(ctx, &user.RegisterRequest{
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
	// connect to gRPC
	conn, err := connectToUserService()
	if err != nil {
		go app.logger.LogFatal(&logs.Log{
			Service: "user-service",
			Message: "Error on connecting to the user-service",
			Error:   err.Error(),
		})
		app.errorJSON(w, errors.New("service is unavailable. Please try again later"))
		return
	}
	defer conn.Close()

	// create client
	client := user.NewAuthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := client.Login(ctx, &user.AuthRequest{
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
