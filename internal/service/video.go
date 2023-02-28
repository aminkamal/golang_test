package service

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateVideoRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	URL         string `json:"url" binding:"required"`
}

type Video struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Duration    int       `json:"duration"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (svc *Service) HandleGetVideos(c *gin.Context) {
	var videos []Video

	svc.DB.Find(&videos)

	c.JSON(http.StatusOK, videos)
}

func (svc *Service) HandleCreateVideo(c *gin.Context) {
	var req CreateVideoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		WriteErrorResponse(c, err)
		return
	}

	// Assume the duration (in seconds) of the video came from another service
	// using the URL in the request (e.g. after it has been transcoded or processed)
	duration := 3600

	video := Video{
		Name:        req.Name,
		Description: req.Description,
		URL:         req.URL,
		Duration:    duration,
	}

	if result := svc.DB.Create(&video); result.Error != nil {
		WriteErrorResponse(c, ErrInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, video)
}

func (svc *Service) HandleGetVideo(c *gin.Context) {
	video, err := svc.getVideoByID(c)
	if err != nil {
		WriteErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, video)
}

func (svc *Service) HandleDeleteVideo(c *gin.Context) {
	video, err := svc.getVideoByID(c)
	if err != nil {
		WriteErrorResponse(c, err)
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
