package service

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func New() *Service {
	db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/video_db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Service{
		DB: db,
	}
}
