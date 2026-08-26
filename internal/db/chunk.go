package db

import (
	"context"
	"polaris/internal/config"

	"github.com/pgvector/pgvector-go"
)

func (s *Service) CreateChunk(ctx context.Context, sourceID uint, text string, vector pgvector.Vector, contentType ContentType) (*Chunk, error) {
	chunk := &Chunk{
		SourceID: sourceID,
		Text:     text,
		Vector:   vector,
		Type:     contentType,
	}

	err := s.db.WithContext(ctx).Create(chunk).Error
	if err != nil {
		return nil, err
	}

	return chunk, nil
}

type ChunkWithScore struct {
	Chunk
	Distance float32
}

func (s *Service) FindChunks(ctx context.Context, queryVector []float32) ([]ChunkWithScore, error) {
	var results []ChunkWithScore
	vector := pgvector.NewVector(queryVector)

	err := s.db.
		WithContext(ctx).
		Model(&Chunk{}).
		Select("*, vector <=> ? AS distance", vector).
		Order("distance").
		Limit(config.ChunksPerRequest).
		Scan(&results).Error

	return results, err
}
