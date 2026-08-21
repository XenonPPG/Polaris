package pipeline

import (
	"polaris/internal/db"
	"polaris/internal/embedding"
)

type Service struct {
	embeddingService *embedding.Service
	dbService        *db.Service
}

func New(embeddingService *embedding.Service, dbService *db.Service) *Service {
	s := &Service{
		embeddingService: embeddingService,
		dbService:        dbService,
	}

	return s
}
