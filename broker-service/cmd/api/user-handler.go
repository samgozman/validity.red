package main

import (
	"context"
	"net/http"
	"time"

	"github.com/samgozman/validity.red/broker/proto/logs"
	"github.com/samgozman/validity.red/broker/proto/user"
)

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