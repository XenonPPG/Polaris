package db

import (
	"context"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model

	FromAI  bool   `gorm:"not_null"`
	Content string `gorm:"not_null"`
	ChatID  uint
	Chat    Chat
}

func (s *Service) CreateMessage(ctx context.Context, chatID uint, content string, fromAI bool) (*Message, error) {
	message := &Message{
		Content: content,
		ChatID:  chatID,
		FromAI:  fromAI,
	}

	err := s.db.WithContext(ctx).Create(message).Error
	if err != nil {
		return nil, err
	}

	return message, nil
}
