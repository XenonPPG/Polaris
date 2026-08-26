package db

import (
	"github.com/pgvector/pgvector-go"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model

	Name      string      `gorm:"not null;check:char_length(name) BETWEEN 1 AND 60"`
	Data      []byte      `gorm:"not null"`
	Type      ContentType `gorm:"not null"`
	SizeBytes int         `gorm:"not null"`
}

type ContentType string

const (
	Text ContentType = "text"
	File ContentType = "file"
)

type Chunk struct {
	gorm.Model

	SourceID uint            `gorm:"not null"`
	Text     string          `gorm:"not null"`
	Vectors  pgvector.Vector `gorm:"type:vector(768);not null"`
	Type     ContentType     `gorm:"not null"`
}
