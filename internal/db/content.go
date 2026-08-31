package db

import (
	"context"
	"polaris/internal/types"

	"gorm.io/gorm"
)

type Content struct {
	gorm.Model

	Name      string            `gorm:"not null;check:char_length(name) BETWEEN 1 AND 60"`
	Data      []byte            `gorm:"not null"`
	Type      types.ContentType `gorm:"not null"`
	SizeBytes int               `gorm:"not null"`
}

func (s *Service) CreateContent(ctx context.Context, name string, data []byte, contentType types.ContentType) (*Content, error) {
	content := &Content{
		Name:      name,
		Data:      data,
		Type:      contentType,
		SizeBytes: len(data),
	}

	err := s.db.WithContext(ctx).Create(content).Error
	if err != nil {
		return nil, err
	}

	return content, nil
}
