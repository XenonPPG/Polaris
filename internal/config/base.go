package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	PostgresUser     string `env:"POSTGRES_USER" env-default:"postgres"`
	PostgresPassword string `env:"POSTGRES_PASSWORD" env-default:"5432"`
	PostgresHost     string `env:"POSTGRES_HOST" env-required:"true"`
	PostgresPort     string `env:"POSTGRES_PORT" env-required:"true"`
	PostgresDB       string `env:"POSTGRES_DB" env-required:"true"`

	EmbeddingURL string `env:"EMBEDDING_URL" env-required:"true"`
}

const (
	SimilarityThreshold = 0.3
	ChunksPerRequest    = 3
)

func LoadConfig(path string) (config Config, err error) {
	err = cleanenv.ReadConfig(path, &config)
	return
}
