package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateAPIKey(c *gin.Context) {
	apiKey := c.GetHeader("Authorization")
	if strings.TrimSpace(apiKey) == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "unauthorized",
		})
		c.Abort()
	}
	c.Next()
}
