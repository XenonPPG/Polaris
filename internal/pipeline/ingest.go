package pipeline

import (
	"context"
	"math"
	"polaris/internal/config"
	"polaris/internal/pipeline/splitters"

	"github.com/pgvector/pgvector-go"
)

func (s *Service) Ingest(ctx context.Context, content Content) error {
	// save content to db
	instance, err := s.dbService.CreateContent(ctx, content.Name, content.Data, content.Type)
	if err != nil {
		return err
	}

	// vectors
	chunks, err := s.fillVector(ctx, splitters.Split(content.Data, content.Type))
	if err != nil {
		return err
	}

	// merge chunks
	var merged []Chunk
	if len(chunks) > 0 {
		merged = make([]Chunk, 0)
		currentData := chunks[0].Data

		for i := 0; i < len(chunks)-1; i++ {
			similarity := cosineSimilarity(chunks[i].Vector, chunks[i+1].Vector)
			//log.Printf("text 1: %s\ntext 2:%s\nsimilarity:%.2f\n\n", chunks[i].Data, chunks[i+1].Data, similarity)
			if similarity >= config.SimilarityThreshold {
				currentData += chunks[i+1].Data
				continue
			}
			merged = append(merged, Chunk{Data: currentData})
			currentData = chunks[i+1].Data
		}
		merged = append(merged, Chunk{Data: currentData})
	}

	// recalculate vector
	mergedParts := make([]string, 0, len(merged))
	for _, c := range merged {
		mergedParts = append(mergedParts, c.Data)
	}

	mergedWithVector, err := s.fillVector(ctx, mergedParts)
	if err != nil {
		return err
	}

	// save to db
	for _, c := range mergedWithVector {
		if _, err := s.dbService.CreateChunk(
			ctx,
			instance.ID,
			c.Data,
			pgvector.NewVector(c.Vector),
			content.Type,
		); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) fillVector(ctx context.Context, parts []string) ([]Chunk, error) {
	chunks := make([]Chunk, 0, len(parts))
	for _, p := range parts {
		vec, err := s.embeddingService.Embed(ctx, p)
		if err != nil {
			return nil, err
		}
		if vec == nil {
			continue
		}
		chunks = append(chunks, Chunk{
			Data:   p,
			Vector: vec,
		})
	}
	return chunks, nil
}

func cosineSimilarity(a, b []float32) float32 {
	var dot, normA, normB float32
	for i := range a {
		dot += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return dot / (float32(math.Sqrt(float64(normA))) * float32(math.Sqrt(float64(normB))))
}
