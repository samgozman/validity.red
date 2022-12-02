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
	res, err := app.usersClient.userService.Register(ctx, &user.RegisterRequest{
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

	// Create verification token for the user to verify email
	verificationToken, err := app.token.Generate(res.UserId, app.options.JWTVerificationTTL)
	if err != nil {
		log.Println("Error on calling user-service::Register::Generate token method:", err)
		c.Error(err)
		return
	}
	// Save verification token to Redis with 24h TTL
	app.redisClient.SetNX(
		ctx,
		"user:verification:"+res.UserId, verificationToken,
		time.Second*time.Duration(app.options.JWTVerificationTTL),
	)

	if app.options.Environment == "production" {
		verificationLink := app.options.AppUrl + "/verify/" + verificationToken
		app.mailer.SendEmailVerification(requestPayload.Email, verificationLink)
	}

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
	token, err := app.token.Generate(res.UserId, app.options.JWTAuthTTL)
	if err != nil {
		log.Println("Error on calling gateway-service::token::Generate method:", err)
		c.Error(err)
		return
	}

	// write jwt token
	c.SetCookie("token", token, app.options.JWTAuthTTL, "/", "", false, false)
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
	token, err := app.token.Refresh(tk.(string), app.options.JWTAuthTTL)
	if err != nil {
		log.Println("Error on calling gateway-service::token::Refresh method:", err)
		c.Error(ErrUnauthorized)
		return
	}

	c.SetCookie("token", token, app.options.JWTAuthTTL, "/", "", false, false)
	c.Status(http.StatusAccepted)
}
