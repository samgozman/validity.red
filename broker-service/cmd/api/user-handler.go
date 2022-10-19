package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samgozman/validity.red/broker/proto/user"
)

type AuthPayload struct {
	Email    string `json:"email" uri:"email" binding:"required,email"`
	Password string `json:"password" uri:"password" binding:"required,min=8,max=64"`
}

// Call Register method on `user-service`
func (app *Config) userRegister(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var payload jsonResponse
	requestPayload := AuthPayload{}
	if err := c.BindJSON(&requestPayload); err != nil {
		c.Error(ErrInvalidInputs)
		return
	}

	// call service
	res, err := app.usersClient.userService.Register(ctx, &user.RegisterRequest{
		RegisterEntry: &user.Register{
			Email:    requestPayload.Email,
			Password: requestPayload.Password,
		},
	})
	if err != nil {
		log.Println("Error on calling user-service::Register method:", err)
		c.Error(err)
		return
	}

	// TODO: Send verification email

	payload.Error = false
	payload.Message = res.Result

	c.JSON(http.StatusCreated, payload)
}

// Call Login method on `user-service`
func (app *Config) userLogin(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var payload jsonResponse
	requestPayload := AuthPayload{}
	if err := c.BindJSON(&requestPayload); err != nil {
		c.Error(ErrInvalidInputs)
		return
	}

	// call service
	res, err := app.usersClient.authService.Login(ctx, &user.AuthRequest{
		AuthEntry: &user.Auth{
			Email:    requestPayload.Email,
			Password: requestPayload.Password,
		},
	})
	if err != nil {
		log.Println("Error on calling user-service::Login method:", err)
		c.Error(err)
		return
	}

	payload.Error = false
	payload.Message = res.Result
	payload.Data = struct {
		CalendarId string `json:"calendarId"`
	}{
		CalendarId: res.CalendarId,
	}

	// Generate JWT token
	token, err := app.token.Generate(res.UserId)
	if err != nil {
		log.Println("Error on calling broker-service::token::Generate method:", err)
		c.Error(err)
		return
	}

	// write jwt token
	c.SetCookie("token", token, app.token.MaxAge, "/", "", false, false)

	c.JSON(http.StatusAccepted, payload)
}

// Refresh current user JWT token
func (app *Config) userRefreshToken(c *gin.Context) {
	// get token from context
	tk, _ := c.Get("Token")

	var payload jsonResponse

	// Refresh JWT token
	token, err := app.token.Refresh(tk.(string))
	if err != nil {
		log.Println("Error on calling broker-service::token::Refresh method:", err)
		c.Error(ErrUnauthorized)
		return
	}

	c.SetCookie("token", token, app.token.MaxAge, "/", "", false, false)

	payload.Error = false
	payload.Message = "Token refreshed"

	c.JSON(http.StatusAccepted, payload)
}
