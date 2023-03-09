package middleware

import (
	"strings"

	"github.com/aminkamal/golang_test/pkg/response"
	"github.com/gin-gonic/gin"
)

func ValidateAPIKey(c *gin.Context) {
	apiKey := c.GetHeader("Authorization")
	if strings.TrimSpace(apiKey) == "" {
		response.WriteErrorResponse(c, response.ErrMissingAPIKey)
		c.Abort()
	}
	if apiKey != "hunter2" {
		response.WriteErrorResponse(c, response.ErrInvalidAPIKey)
		c.Abort()
	}

	c.Next()
}
