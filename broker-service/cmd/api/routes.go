package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *Config) routes() *gin.Engine {
	g := gin.Default()

	g.Use(cors.New(cors.Config{
		// TODO: Set to validity.red domains
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
		AllowWildcard:    true,
	}))

	documents := g.Group("/documents")
	documents.Use(app.AuthGuard())
	{
		documents.GET("", app.documentGetAll)
		documents.GET("/:documentId", app.documentGetOne)
		documents.GET("/:documentId/notifications", app.documentNotificationGetAll)
		documents.POST("/:documentId/notifications/create", app.documentNotificationCreate)
		documents.DELETE("/:documentId/notifications/delete/:id", app.documentNotificationDelete)
		documents.PATCH("/:documentId/notifications/edit/:id", app.documentNotificationEdit)
		documents.POST("/create", app.documentCreate)
		// TODO: edit/:id
		documents.PATCH("/edit", app.documentEdit)
		documents.DELETE("/:documentId/delete", app.documentDelete)
		documents.GET("/statistics", app.documentGetStatistics)
	}

	calendar := g.Group("/calendar")
	calendar.Use(app.AuthGuard())
	{
		calendar.GET("", app.getCalendar)
	}

	user := g.Group("/user")
	user.Use(app.AuthGuard())
	{
		user.GET("/token/refresh", app.userRefreshToken)
	}

	// Auth routes (without auth guard)
	auth := g.Group("/auth")
	{
		auth.POST("/login", app.userLogin)
		auth.POST("/register", app.userRegister)
	}

	return g
}
