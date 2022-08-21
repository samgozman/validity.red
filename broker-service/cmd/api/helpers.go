package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Decode JSON request to any struct
func decodeJSON[T any](c *gin.Context) T {
	var requestPayload T

	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return requestPayload
	}

	return requestPayload
}
