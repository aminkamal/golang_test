package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aminkamal/golang_test/internal/service"
	"github.com/stretchr/testify/assert"
)

const (
	validAPIKey = "hunter2"
)

func TestAuthMissingAPIKey(t *testing.T) {
	svc := service.New()
	svc.AddRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/videos", nil)
	svc.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestAuthSuccess(t *testing.T) {
	svc := service.New()
	svc.AddRoutes()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/videos", nil)
	req.Header.Add("Authorization", validAPIKey)
	svc.Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
