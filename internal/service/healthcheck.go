package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (svc *Service) HandleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
