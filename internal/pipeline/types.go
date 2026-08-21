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
	SourceID string
	Data     string
	Vectors  embedding.Vector
	Type     db.ContentType
}
