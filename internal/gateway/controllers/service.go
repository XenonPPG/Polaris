package controllers

import (
	"polaris/internal/db"
	"polaris/internal/pipeline"
)

type Service struct {
	mainService *pipeline.Service
	dbService   *db.Service
}

func NewService(mainService *pipeline.Service, dbService *db.Service) *Service {
	return &Service{
		mainService: mainService,
		dbService:   dbService,
	}
}
