package db

import "context"

func (s *Service) CreateContent(ctx context.Context, name string, data []byte, contentType ContentType) (*Content, error) {
	content := &Content{
		Name:      name,
		Data:      data,
		Type:      contentType,
		SizeBytes: len(data),
	}

	_, err := s.db.NewInsert().Model(content).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *Service) GetContent(ctx context.Context, id string) (*Content, error) {
	var content *Content

	_, err := s.db.NewSelect().Model(content).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (s *Service) DeleteContent(ctx context.Context, id string) error {
	_, err := s.db.NewDelete().Model(&Content{}).Where("id = ?", id).Exec(ctx)
	if err != nil {
		return err
	}
	err = s.DeleteChunks(ctx, id)

	return err
}
