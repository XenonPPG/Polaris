package db

import "context"

func (s *Service) CreateContent(ctx context.Context, name string, data []byte, contentType ContentType) (*Content, error) {
	content := &Content{
		Name:      name,
		Data:      data,
		Type:      contentType,
		SizeBytes: len(data),
	}

	err := s.db.WithContext(ctx).Create(content).Error
	if err != nil {
		return nil, err
	}

	return content, nil
}
