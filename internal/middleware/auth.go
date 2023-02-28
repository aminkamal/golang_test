package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func writeError(c *gin.Context, err string) {
	c.JSON(http.StatusForbidden, gin.H{
		"error": err,
	})
}

func ValidateAPIKey(c *gin.Context) {
	apiKey := c.GetHeader("Authorization")
	if strings.TrimSpace(apiKey) == "" {
		writeError(c, "missing api key")
		c.Abort()
	}

	if apiKey != "hunter2" {
		writeError(c, "invalid api key")
		c.Abort()
	}

	c.Next()
}
