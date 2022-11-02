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

type RegisterPayload struct {
	Email    string `json:"email" uri:"email" binding:"required,email"`
	Password string `json:"password" uri:"password" binding:"required,min=8,max=64"`
	Timezone string `json:"timezone" uri:"timezone" binding:"required,timezone"`
}

// Call Register method on `user-service`
func (app *Config) userRegister(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	requestPayload := RegisterPayload{}
	if err := c.BindJSON(&requestPayload); err != nil {
		c.Error(ErrInvalidInputs)
		return
	}

	// call service
	_, err := app.usersClient.userService.Register(ctx, &user.RegisterRequest{
		RegisterEntry: &user.Register{
			Email:    requestPayload.Email,
			Password: requestPayload.Password,
			Timezone: requestPayload.Timezone,
		},
	})
	if err != nil {
		log.Println("Error on calling user-service::Register method:", err)
		c.Error(err)
		return
	}

	// TODO: Send verification email

	c.Status(http.StatusCreated)
}

// Call Login method on `user-service`
func (app *Config) userLogin(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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

	// Generate JWT token
	token, err := app.token.Generate(res.UserId)
	if err != nil {
		log.Println("Error on calling gateway-service::token::Generate method:", err)
		c.Error(err)
		return
	}

	// write jwt token
	c.SetCookie("token", token, app.token.MaxAge, "/", "", false, false)
	c.JSON(http.StatusAccepted, struct {
		CalendarId string `json:"calendarId"`
		Timezone   string `json:"timezone"`
	}{
		CalendarId: res.CalendarId,
		Timezone:   res.Timezone,
	})
}

// Refresh current user JWT token
func (app *Config) userRefreshToken(c *gin.Context) {
	// get token from context
	tk, _ := c.Get("Token")

	// Refresh JWT token
	token, err := app.token.Refresh(tk.(string))
	if err != nil {
		log.Println("Error on calling gateway-service::token::Refresh method:", err)
		c.Error(ErrUnauthorized)
		return
	}

	c.SetCookie("token", token, app.token.MaxAge, "/", "", false, false)
	c.Status(http.StatusAccepted)
}
