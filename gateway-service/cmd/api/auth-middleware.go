package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware that checks if the user is authenticated
// and passes the UserId to the next handler via context
func (app *Config) AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			userId string
			token  string
		)

		// Get token from cookie
		token, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Verify token and decode UserId from it
		userId, err = app.token.Verify(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Add decoded user id for the context
		c.Set("UserId", userId)
		c.Set("Token", token)

		c.Next()
	}
}
