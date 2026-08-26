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

func (s *Service) GetContent(ctx context.Context, id uint) (*Content, error) {
	var content *Content

	err := s.db.WithContext(ctx).Where("id = ?", id).First(&content).Error
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *Service) DeleteContent(ctx context.Context, id uint) error {
	err := s.db.WithContext(ctx).Where("id = ?", id).Delete(&Content{}).Error
	if err != nil {
		return err
	}
	err = s.DeleteChunks(ctx, id)

	return err
}
