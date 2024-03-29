package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/samgozman/validity.red/broker/proto/user"
)

type authPayload struct {
	Email    string `json:"email" uri:"email" binding:"required,email"`
	Password string `json:"password" uri:"password" binding:"required,min=8,max=64"`
}

type registerPayload struct {
	Email            string `json:"email" uri:"email" binding:"required,email"`
	Password         string `json:"password" uri:"password" binding:"required,min=8,max=64"`
	Timezone         string `json:"timezone" uri:"timezone" binding:"required,timezone"`
	HCaptchaResponse string `json:"hcaptcha" uri:"hcaptcha" binding:"required"`
}

type emailVerificationPayload struct {
	Token string `json:"token" uri:"token" binding:"required,jwt"`
}

// Call Register method on `user-service`.
func (app *Config) userRegister(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	requestPayload := registerPayload{}
	if err := c.BindJSON(&requestPayload); err != nil {
		_ = c.Error(ErrInvalidInputs)
		return
	}

	// If environment is not production, skip captcha verification
	if app.options.Environment == "production" {
		if hr := app.hcaptcha.VerifyToken(requestPayload.HCaptchaResponse); !hr.Success {
			sentry.CaptureException(fmt.Errorf("hCaptcha errors: %s", hr.ErrorCodes))
			_ = c.Error(ErrInvalidCaptcha)
			return
		}
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
		_ = c.Error(err)

		return
	}

	// Create verification token for the user to verify email
	verificationToken, err := app.token.Generate(res.UserId, app.options.JWTVerificationTTL)
	if err != nil {
		log.Println("Error on calling user-service::Register::Generate token method:", err)
		_ = c.Error(err)

		return
	}
	// Save verification token to Redis with 24h TTL
	app.redisClient.SetNX(
		ctx,
		"user:verification:"+res.UserId, verificationToken,
		time.Second*time.Duration(app.options.JWTVerificationTTL),
	)

	if app.options.Environment == "production" {
		verificationLink := app.options.AppURL + "/verify?token=" + verificationToken
		err := app.mailer.SendEmailVerification(requestPayload.Email, verificationLink)

		if err != nil {
			sentry.CaptureException(fmt.Errorf("SendEmailVerification error: %w", err))
		}
	}

	c.Status(http.StatusCreated)
}

// Call Login method on `user-service`.
func (app *Config) userLogin(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	requestPayload := authPayload{}
	if err := c.BindJSON(&requestPayload); err != nil {
		_ = c.Error(ErrInvalidInputs)
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
		_ = c.Error(err)

		return
	}

	// User's email should be verified before login
	if !res.IsVerified {
		_ = c.Error(ErrEmailNotVerified)
		return
	}

	// Generate JWT token
	token, err := app.token.Generate(res.UserId, app.options.JWTAuthTTL)
	if err != nil {
		log.Println("Error on calling gateway-service::token::Generate method:", err)
		_ = c.Error(err)

		return
	}

	// write jwt token
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", token, app.options.JWTAuthTTL, "/", "", true, false)
	c.JSON(http.StatusAccepted, struct {
		CalendarID string `json:"calendarId"`
		Timezone   string `json:"timezone"`
	}{
		CalendarID: res.CalendarId,
		Timezone:   res.Timezone,
	})
}

// Refresh current user JWT token.
func (app *Config) userRefreshToken(c *gin.Context) {
	// get token from context
	tk, _ := c.Get("Token")

	// Refresh JWT token
	token, err := app.token.Refresh(tk.(string), app.options.JWTAuthTTL)
	if err != nil {
		log.Println("Error on calling gateway-service::token::Refresh method:", err)

		_ = c.Error(ErrUnauthorized)

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", token, app.options.JWTAuthTTL, "/", "", false, false)
	c.Status(http.StatusAccepted)
}

// Verify user email by sended token.
func (app *Config) userVerifyEmail(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	json := emailVerificationPayload{}

	// Validate inputs
	if err := c.BindJSON(&json); err != nil {
		_ = c.Error(ErrInvalidInputs)
		return
	}

	// Validate token
	userID, err := app.token.Verify(json.Token)
	if err != nil {
		_ = c.Error(ErrUnauthorized)
		return
	}

	// Check if token exists in Redis for this user
	token, err := app.redisClient.Get(ctx, "user:verification:"+userID).Result()
	if err != nil {
		_ = c.Error(ErrUnauthorized)
		return
	}

	// Check if token is valid
	if token != json.Token {
		_ = c.Error(ErrUnauthorized)
		return
	}

	// Delete token from Redis
	app.redisClient.Del(ctx, "user:verification:"+userID)

	// Call user-service to verify user
	_, err = app.usersClient.userService.SetIsVerified(ctx, &user.SetIsVerifiedRequest{
		UserId:     userID,
		IsVerified: true,
	})
	if err != nil {
		log.Println("Error on calling user-service::SetIsVerified method:", err)
		_ = c.Error(err)

		return
	}

	c.Status(http.StatusAccepted)
}
