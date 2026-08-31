package pipeline

import (
	"polaris/internal/embedding"
	"polaris/internal/types"
)

type Content struct {
	Name string
	Data []byte
	Type types.ContentType
}

type Chunk struct {
	SourceID uint
	Data     string
	Vector   embedding.Vector
	Type     types.ContentType
}

type ChunkWithScore struct {
	Chunk
	Distance float32
}

func (c ChunkWithScore) GetSimilarity() float32 {
	return 1 - c.Distance
}
