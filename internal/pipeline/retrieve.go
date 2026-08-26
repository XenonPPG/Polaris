package pipeline

import (
	"context"
	"errors"
)

func (s *Service) Retrieve(ctx context.Context, query string) ([]ChunkWithScore, error) {
	queryVector, err := s.embeddingService.Embed(ctx, query)
	if err != nil {
		return nil, err
	}
	if queryVector == nil || len(queryVector) == 0 {
		return nil, errors.New("query vector is empty")
	}

	results, err := s.dbService.FindChunks(ctx, queryVector)
	if err != nil {
		return nil, err
	}

	chunksWithScore := make([]ChunkWithScore, 0, len(results))
	for _, r := range results {
		chunksWithScore = append(chunksWithScore, ChunkWithScore{
			Chunk: Chunk{
				SourceID: r.SourceID,
				Data:     r.Text,
				Vectors:  r.Vectors.Slice(),
				Type:     r.Type,
			},
			Distance: r.Distance,
		})
	}

	return chunksWithScore, nil
}
