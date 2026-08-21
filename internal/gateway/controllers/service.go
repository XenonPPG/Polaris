package controllers

import (
	"polaris/internal/pipeline"
)

type Service struct {
	mainService *pipeline.Service
}

func NewService(mainService *pipeline.Service) *Service {
	return &Service{
		mainService: mainService,
	}
}
