package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

// A middleware that is used to handle errors.
func (app *Config) ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		for _, err := range c.Errors {
			unwrappedErr := err.Unwrap()

			// If error is of type from defined errors
			for _, errType := range ErrorsArr {
				if !errors.Is(unwrappedErr, errType) {
					continue
				}
				c.AbortWithStatusJSON(ErrorStatus[errType], gin.H{
					"error":   true,
					"message": errType.Error(),
				})
				return
			}

			// If error is of type from gRPC
			if st, ok := status.FromError(unwrappedErr); ok {
				c.AbortWithStatusJSON(RPCStatus[st.Code()], gin.H{
					"error":   true,
					"message": st.Message(),
				})
				return
			}
		}

		// TODO: Send all internal errors to Sentry

		// if not returned yet => return 500
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
