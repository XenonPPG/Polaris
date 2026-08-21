package embedding

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Service struct {
	embeddingURL string
	client       *http.Client
}

func New(embeddingURL string, timeout time.Duration) *Service {
	return &Service{
		embeddingURL: embeddingURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (s *Service) Embed(ctx context.Context, content string) (Vector, error) {
	payload := Payload{
		Inputs: content,
	}
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.embeddingURL, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected status code: " + resp.Status)
	}

	var vectors []Vector
	err = json.NewDecoder(resp.Body).Decode(&vectors)
	if err != nil {
		return nil, err
	}

	return vectors[0], nil
}
