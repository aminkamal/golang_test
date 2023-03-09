package service

import (
	"github.com/aminkamal/golang_test/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func New() *Service {
	// TODO: read from .env file
	db, err := gorm.Open(postgres.Open("postgres://postgres:password@db:5432/video_db"), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/video_db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()

	return &Service{
		DB:     db,
		Router: router,
	}
}

func (svc *Service) AddRoutes() {

	svc.Router.GET("/healthcheck", svc.HandleHealthCheck)

	apiV1 := svc.Router.Group("/v1", middleware.ValidateAPIKey)
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
}
