package pipeline

import (
	"polaris/internal/db"
	"polaris/internal/embedding"
)

type Content struct {
	Name string
	Data []byte
	Type db.ContentType
}

type Chunk struct {
	SourceID uint
	Data     string
	Vector   embedding.Vector
	Type     db.ContentType
}

type ChunkWithScore struct {
	Chunk
	Distance float32
}

func (c ChunkWithScore) GetSimilarity() float32 {
	return 1 - c.Distance
}
