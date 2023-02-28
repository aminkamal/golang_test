package service

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidUUID         = errors.New("invalid id, not a valid uuid")
	ErrInternalServerError = errors.New("internal server error")
	ErrResourceNotFound    = errors.New("resource was not found")
)

func WriteErrorResponse(c *gin.Context, err error) {
	if errors.Is(err, ErrInvalidUUID) {
		writeError(c, http.StatusBadRequest, err)
	} else if errors.Is(err, ErrInternalServerError) {
		// TODO: Log this error
		writeError(c, http.StatusInternalServerError, err)
	} else if errors.Is(err, ErrResourceNotFound) {
		writeError(c, http.StatusNotFound, err)
	} else {
		writeError(c, http.StatusBadRequest, err)
	}
}

func writeError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}
