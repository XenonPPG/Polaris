package db

import (
	"context"
	"polaris/internal/gateway/utils"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model

	Title    string
	Messages []Message
}

func (s *Service) CreateChat(ctx context.Context) (*Chat, error) {
	chat := &Chat{}

	err := s.db.WithContext(ctx).Create(chat).Error
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (s *Service) GetChat(ctx context.Context, id uint) (*Chat, error) {
	chat := &Chat{}
	err := s.db.
		WithContext(ctx).
		Preload("Messages").
		First(chat, id).Error
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (s *Service) ListChats(ctx context.Context, limit, offset int) ([]Chat, error) {
	chats := make([]Chat, 0)

	const maxPerRequest = 100
	if limit < 0 {
		limit = maxPerRequest
	}
	limit = utils.Clamp(limit, 1, maxPerRequest)
	offset = max(0, offset)
	err := s.db.
		WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&chats).Error
	if err != nil {
		return nil, err
	}

	return chats, nil
}
