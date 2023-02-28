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
		apiV1.POST("/videos", svc.HandleCreateVideo)
		apiV1.GET("/videos/:videoID", svc.HandleGetVideo)
		apiV1.DELETE("/videos/:videoID", svc.HandleDeleteVideo)

		apiV1.POST("/videos/:videoID/annotations", svc.HandleCreateAnnotation)
		apiV1.GET("/videos/:videoID/annotations", svc.HandleGetAnnotations)
		apiV1.PUT("/videos/:videoID/annotations/:annotationID", svc.HandlePutAnnotation)
		apiV1.DELETE("/videos/:videoID/annotations/:annotationID", svc.HandleDeleteAnnotation)
		apiV1.GET("/videos/:videoID/annotations/:annotationID", svc.HandleGetAnnotation)
	}

	r.Run()
}
