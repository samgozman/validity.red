package main

import (
	"errors"
	"net/http"
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

// authRegister temprary stub
func (app *Config) authRegister(w http.ResponseWriter, authPayload AuthPayload) {
	app.writeJSON(w, http.StatusAccepted, map[string]string{"message": "registration successful"})
}
