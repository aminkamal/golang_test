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

func writeErrorResponse(c *gin.Context, err error) {
	if errors.Is(err, ErrInvalidUUID) {
		c.JSON(http.StatusBadRequest, nil)
	}

	if errors.Is(err, ErrInternalServerError) {
		// TODO: Log this error
		c.JSON(http.StatusInternalServerError, nil)
	}

	if errors.Is(err, ErrInternalServerError) {
		c.JSON(http.StatusNotFound, nil)
	}
}
