package config

type Config struct {
	EmbeddingURL string
}

const (
	SimilarityThreshold = 0.8
	ChunksPerRequest    = 3
)
