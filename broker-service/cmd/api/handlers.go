package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/samgozman/validity.red/broker/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
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
	case "AuthRegister":
		app.authRegister(w, requestPayload.Auth)
	default:
		app.errorJSON(w, errors.New("invalid action"))
	}
}

// Call Register method on `user-service`
func (app *Config) authRegister(w http.ResponseWriter, authPayload AuthPayload) {
	// connect to gRPC
	authURL := fmt.Sprintf("user-service:%s", os.Getenv("USER_GRPC_PORT"))
	conn, err := grpc.Dial(authURL, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		app.errorJSON(w, err)
	}
	defer conn.Close()

	// create client
	client := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// call service
	res, err := client.Register(ctx, &user.RegisterRequest{
		RegisterEntry: &user.Register{
			Email:    authPayload.Email,
			Password: authPayload.Password,
		},
	})
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = res.Result

	app.writeJSON(w, http.StatusAccepted, payload)
}
