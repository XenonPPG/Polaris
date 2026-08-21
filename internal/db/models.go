package db

import (
	"github.com/pgvector/pgvector-go"
	"github.com/uptrace/bun"
)

type Content struct {
	bun.BaseModel `bun:"table:contents"`

	ID        string      `bun:"id,pk"`
	Name      string      `bun:"name,notnull"`
	Data      []byte      `bun:"data,notnull,type:bytea"`
	Type      ContentType `bun:"type,type:ContentType,notnull"`
	SizeBytes int         `bun:"size_bytes,notnull"`
}

type ContentType string

const (
	Text ContentType = "text"
	File ContentType = "file"
)

type Chunk struct {
	bun.BaseModel `bun:"table:chunks"`

	ID       string          `bun:"id,pk,default:gen_random_uuid()"`
	SourceID string          `bun:"source_id,notnull"`
	Text     string          `bun:"text,notnull"`
	Vectors  pgvector.Vector `bun:"embedding,type:vector(768)"`
	Type     ContentType     `bun:"type,type:ContentType,notnull"`
}
