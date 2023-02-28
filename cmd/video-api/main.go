package main

import (
	"github.com/aminkamal/golang_test/internal/middleware"
	"github.com/aminkamal/golang_test/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	svc := service.New()

	r.GET("/healthcheck", svc.HandleHealthCheck)

	apiV1 := r.Group("/v1", middleware.ValidateAPIKey)
	{
		apiV1.GET("/videos", svc.HandleGetVideos)
		apiV1.GET("/videos/:videoID", svc.HandleGetVideo)
	}

	r.Run()
}
