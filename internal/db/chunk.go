package db

import (
	"context"

	"github.com/pgvector/pgvector-go"
)

func (s *Service) CreateChunk(ctx context.Context, sourceID uint, text string, vectors pgvector.Vector, contentType ContentType) (*Chunk, error) {
	chunk := &Chunk{
		SourceID: sourceID,
		Text:     text,
		Vectors:  vectors,
		Type:     contentType,
	}

	err := s.db.WithContext(ctx).Create(chunk).Error
	if err != nil {
		return nil, err
	}

	return chunk, nil
}

func (s *Service) GetChunk(ctx context.Context, id uint) (*Chunk, error) {
	var chunk *Chunk

	err := s.db.WithContext(ctx).Where("id = ?", id).First(&chunk).Error
	if err != nil {
		return nil, err
	}

	return chunk, nil
}

func (s *Service) GetChunks(ctx context.Context, sourceID uint) ([]Chunk, error) {
	chunks := make([]Chunk, 0)

	err := s.db.WithContext(ctx).Where("source_id = ?", sourceID).Find(&chunks).Error
	if err != nil {
		return nil, err
	}

	return chunks, nil
}

func (s *Service) DeleteChunk(ctx context.Context, id uint) error {
	err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&Chunk{}).Error
	return err
}

func (s *Service) DeleteChunks(ctx context.Context, sourceID uint) error {
	err := s.db.WithContext(ctx).Where("source_id = ?", sourceID).Delete(&Chunk{}).Error
	return err

}
