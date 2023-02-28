package middleware

import (
	"strings"

	"github.com/aminkamal/golang_test/internal/service"
	"github.com/gin-gonic/gin"
)

func ValidateAPIKey(c *gin.Context) {
	apiKey := c.GetHeader("Authorization")
	if strings.TrimSpace(apiKey) == "" {
		service.WriteErrorResponse(c, service.ErrMissingAPIKey)
		c.Abort()
	}
	if apiKey != "hunter2" {
		service.WriteErrorResponse(c, service.ErrInvalidAPIKey)
		c.Abort()
	}

	c.Next()
}
