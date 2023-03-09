package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrMissingAPIKey       = errors.New("missing api key")
	ErrInvalidAPIKey       = errors.New("invalid api key")
	ErrInvalidUUID         = errors.New("invalid id, not a valid uuid")
	ErrInternalServerError = errors.New("internal server error")
	ErrResourceNotFound    = errors.New("resource was not found")
	ErrInvalidTimeFormat   = errors.New("invalid time format specified")
	ErrInvalidTimeRange    = errors.New("invalid time range specified")
)

func WriteErrorResponse(c *gin.Context, err error) {
	if errors.Is(err, ErrMissingAPIKey) ||
		errors.Is(err, ErrInvalidAPIKey) {
		writeError(c, http.StatusForbidden, err)
	} else if errors.Is(err, ErrInternalServerError) {
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
