package db

import (
	"context"

	"github.com/pgvector/pgvector-go"
)

func (s *Service) CreateChunk(ctx context.Context, sourceID, text string, vectors pgvector.Vector, contentType ContentType) (*Chunk, error) {
	chunk := &Chunk{
		SourceID: sourceID,
		Text:     text,
		Vectors:  vectors,
		Type:     contentType,
	}

	_, err := s.db.NewInsert().Model(chunk).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return chunk, nil
}

func (s *Service) GetChunk(ctx context.Context, id string) (*Chunk, error) {
	var chunk *Chunk

	_, err := s.db.NewSelect().Model(chunk).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return chunk, nil
}

func (s *Service) GetChunks(ctx context.Context, sourceID string) ([]Chunk, error) {
	chunks := make([]Chunk, 0)

	_, err := s.db.NewSelect().Model(chunks).Where("source_id = ?", sourceID).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return chunks, nil
}

func (s *Service) DeleteChunk(ctx context.Context, id string) error {
	_, err := s.db.NewDelete().Model(&Chunk{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (s *Service) DeleteChunks(ctx context.Context, sourceID string) error {
	_, err := s.db.NewDelete().Model(&Chunk{}).Where("source_id = ?", sourceID).Exec(ctx)
	return err

}
