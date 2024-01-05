package main

import (
	"log"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginLimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	storeRedis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func (app *Config) routes() *gin.Engine {
	const requestMaxAge = 300

	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		// TODO: Set to validity.red domains
		AllowOrigins:     []string{"https://validity.host.extr.app/*", "https://validity.red/*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "sentry-trace", "baggage"},
		AllowCredentials: true,
		MaxAge:           requestMaxAge,
		AllowWildcard:    true,
	}))

	engine.Use(sentrygin.New(sentrygin.Options{}))

	g := engine.Group("/api")

	// Rate limiting
	rate, err := limiter.NewRateFromFormatted("1000-H")
	if err != nil {
		log.Fatal(err)
	}
	// Create a store with the redis client.
	store, err := storeRedis.NewStoreWithOptions(app.redisClient, limiter.StoreOptions{
		Prefix: "gin_limiter",
	})
	if err != nil {
		log.Fatal(err)
	}
	// Create a new middleware with the limiter instance.
	rateLimiter := ginLimiter.NewMiddleware(limiter.New(store, rate))
	g.Use(rateLimiter)

	documents := g.Group("/documents")
	documents.Use(app.AuthGuard(), app.ErrorHandler())
	{
		documents.GET("", app.documentGetAll)
		documents.GET("/:documentId", app.documentGetOne)
		documents.GET("/:documentId/notifications", app.documentNotificationGetAll)
		documents.POST("/:documentId/notifications/create", app.documentNotificationCreate)
		documents.DELETE("/:documentId/notifications/delete/:id", app.documentNotificationDelete)
		documents.POST("/create", app.documentCreate)
		documents.PATCH("/edit", app.documentEdit)
		documents.DELETE("/:documentId/delete", app.documentDelete)
		documents.GET("/statistics", app.documentGetStatistics)
	}

	calendar := g.Group("/calendar")
	calendar.Use(app.AuthGuard(), app.ErrorHandler())
	{
		calendar.GET("", app.getCalendar)
	}

	ics := g.Group("/ics")
	ics.Use(app.ErrorHandler())
	{
		ics.GET("/:id", app.getCalendarIcs)
	}

	user := g.Group("/user")
	user.Use(app.AuthGuard(), app.ErrorHandler())
	{
		user.GET("/token/refresh", app.userRefreshToken)
	}

	// Auth routes (without auth guard)
	auth := g.Group("/auth")
	auth.Use(app.ErrorHandler())
	{
		auth.POST("/login", app.userLogin)
		auth.POST("/register", app.userRegister)
		auth.POST("/verify", app.userVerifyEmail)
		// TODO: In separate task
		// auth.POST("/verify/resend", app.userVerifyResend)
	}

	return engine
}
