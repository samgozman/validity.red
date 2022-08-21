package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *Config) routes() *gin.Engine {
	g := gin.Default()

	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		AllowWildcard:    true,
	}))

	handler := g.Group("/handle")
	handler.Use(app.AuthGuard())
	{
		handler.POST("", app.HandleSubmission)
	}

	// Auth routes
	auth := g.Group("/auth")
	{
		auth.POST("/login", app.userLogin)
		auth.POST("/register", app.userRegister)
	}

	return g
}
