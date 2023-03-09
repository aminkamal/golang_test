package service

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aminkamal/golang_test/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateOrUpdateAnnotationRequest struct {
	Type        string `json:"type" binding:"required,oneof=advertisement not_relevant different_language"`
	Note        string `json:"note" binding:"required"`
	StartMarker string `json:"start" binding:"required"`
	EndMarker   string `json:"end" binding:"required"`
}

type Annotation struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()" `
	VideoID     uuid.UUID `json:"-"`
	Type        string    `json:"type"`
	Note        string    `json:"note"`
	StartMarker int       `json:"start"`
	EndMarker   int       `json:"end"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (svc *Service) HandleGetAnnotations(c *gin.Context) {
	video, err := svc.getVideoByID(c)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	var annotations []Annotation

	svc.DB.Find(&annotations, "video_id = ?", video.ID)

	c.JSON(http.StatusOK, annotations)
}

func (svc *Service) HandleCreateAnnotation(c *gin.Context) {
	video, err := svc.getVideoByID(c)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	var req CreateOrUpdateAnnotationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	startMarker, err := parseMarker(req.StartMarker)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	endMarker, err := parseMarker(req.EndMarker)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	if !validRange(*startMarker, *endMarker, video.Duration) {
		response.WriteErrorResponse(c, response.ErrInvalidTimeRange)
		return
	}

	annotation := Annotation{
		VideoID:     video.ID,
		Note:        req.Note,
		Type:        req.Type,
		StartMarker: *startMarker,
		EndMarker:   *endMarker,
	}

	if result := svc.DB.Create(&annotation); result.Error != nil {
		response.WriteErrorResponse(c, response.ErrInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, annotation)
}

func (svc *Service) HandleGetAnnotation(c *gin.Context) {
	annotation, err := svc.getAnnotationByID(c)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, annotation)
}

func (svc *Service) HandlePutAnnotation(c *gin.Context) {
	video, err := svc.getVideoByID(c)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	annotation, err := svc.getAnnotationByID(c)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	var req CreateOrUpdateAnnotationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	startMarker, err := parseMarker(req.StartMarker)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	endMarker, err := parseMarker(req.EndMarker)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	if !validRange(*startMarker, *endMarker, video.Duration) {
		response.WriteErrorResponse(c, response.ErrInvalidTimeRange)
		return
	}

	annotation.Note = req.Note
	annotation.Type = req.Type
	annotation.StartMarker = *startMarker
	annotation.EndMarker = *endMarker

	if result := svc.DB.Save(&annotation); result.Error != nil {
		response.WriteErrorResponse(c, response.ErrInternalServerError)
		return
	}

	c.JSON(http.StatusOK, annotation)
}

func (svc *Service) HandleDeleteAnnotation(c *gin.Context) {
	annotation, err := svc.getAnnotationByID(c)
	if err != nil {
		response.WriteErrorResponse(c, err)
		return
	}

	svc.DB.Delete(&annotation)

	c.Status(http.StatusNoContent)
}

func validRange(startMarker int, endMarker int, duration int) bool {
	return (startMarker) < (endMarker) &&
		(startMarker) <= duration &&
		(endMarker <= duration)
}

func parseMarker(timeStr string) (*int, error) {
	if len(timeStr) != 8 {
		return nil, response.ErrInvalidTimeFormat
	}

	components := strings.Split(timeStr, ":")
	if len(components) != 3 {
		return nil, response.ErrInvalidTimeFormat
	}

	parseComponent := func(numStr string, max int) *int {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil
		}

		if num < 0 || (num > max && max != 0) {
			return nil
		}
		return &num
	}

	hours := parseComponent(components[0], 0)
	minutes := parseComponent(components[1], 59)
	seconds := parseComponent(components[2], 59)

	if hours == nil || minutes == nil || seconds == nil {
		return nil, response.ErrInvalidTimeFormat
	}

	offsetSeconds := (*seconds) + (*minutes)*60 + (*hours)*3600

	return &offsetSeconds, nil
}

func (svc *Service) getAnnotationByID(c *gin.Context) (*Annotation, error) {
	videoID := c.Param("videoID")
	annotationID := c.Param("annotationID")

	_, err := uuid.Parse(videoID)
	if err != nil {
		return nil, response.ErrInvalidUUID
	}

	_, err = uuid.Parse(annotationID)
	if err != nil {
		return nil, response.ErrInvalidUUID
	}

	var annotation Annotation
	if result := svc.DB.First(&annotation, "id = ? and video_id = ?", annotationID, videoID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, response.ErrResourceNotFound
		}

		return nil, response.ErrInternalServerError
	}

	return &annotation, nil
}
