package service_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aminkamal/golang_test/internal/service"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

const (
	validAPIKey = "hunter2"
)

func TestCreateVideo(t *testing.T) {
	svc := service.New()
	svc.AddRoutes()

	createVideoRequest := service.CreateVideoRequest{
		Name:        "Test Video",
		Description: "Not Baby Shark",
		URL:         "https://mybucket.somewhere.com/123/video.mp4",
	}
	request, err := json.Marshal(createVideoRequest)
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/videos", bytes.NewBuffer(request))
	req.Header.Add("Authorization", validAPIKey)
	svc.Router.ServeHTTP(w, req)

	var createVideoResponse service.Video
	err = json.NewDecoder(w.Body).Decode(&createVideoResponse)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, createVideoResponse.Duration, 3600)
	assert.NotNil(t, createVideoResponse.ID)
	assert.NotEmpty(t, createVideoResponse.Name)
	assert.NotEmpty(t, createVideoResponse.Description)
	assert.NotEmpty(t, createVideoResponse.URL)
}
