package service

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Video struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (svc *Service) HandleGetVideos(c *gin.Context) {
	var videos []Video

	svc.DB.Find(&videos)

	c.JSON(http.StatusOK, videos)
}

func (svc *Service) HandleGetVideo(c *gin.Context) {
	video, err := svc.getVideoByID(c)
	if err != nil {
		writeErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, video)
}

func (svc *Service) HandleDeleteVideo(c *gin.Context) {
	video, err := svc.getVideoByID(c)
	if err != nil {
		writeErrorResponse(c, err)
		return
	}

	svc.DB.Delete(&video)

	c.Status(http.StatusNoContent)
}

func (svc *Service) getVideoByID(c *gin.Context) (*Video, error) {
	id := c.Param("videoID")

	_, err := uuid.Parse(id)
	if err != nil {
		return nil, ErrInvalidUUID
	}

	var video Video
	if result := svc.DB.First(&video, "id = ?", id); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrResourceNotFound
		}

		return nil, ErrInternalServerError
	}

	return &video, nil
}
