package main

import (
	"errors"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
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
				status := RPCStatus[st.Code()]
				// Capture all 500 errors to Sentry
				if hub := sentrygin.GetHubFromContext(c); hub != nil && status >= 500 {
					hub.WithScope(func(scope *sentry.Scope) {
						hub.CaptureException(unwrappedErr)
					})
				}

				c.AbortWithStatusJSON(status, gin.H{
					"error":   true,
					"message": st.Message(),
				})

				return
			}
		}

		// if not returned yet => return 500
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.WithScope(func(scope *sentry.Scope) {
				scope.SetExtra("errType", "Partially handled internal error")
				hub.CaptureException(c.Errors[0].Err)
			})
		}

		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
