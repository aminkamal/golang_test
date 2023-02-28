package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Video struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type GetVideosResponse struct {
	Videos []Video `json:"videos"`
}

func (svc *Service) HandleGetVideos(c *gin.Context) {
	videos := GetVideosResponse{
		Videos: make([]Video, 0),
	}
	c.JSON(http.StatusOK, videos)
}

func (svc *Service) HandleGetVideo(c *gin.Context) {
	videoID := c.Param("videoID")

	var video Video
	svc.DB.First(&video, "id = ?", videoID)

	c.JSON(http.StatusOK, video)
}
