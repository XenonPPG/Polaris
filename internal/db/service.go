package db

import (
	"fmt"
	"log"
	"net/url"
	"polaris/internal/config"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func New(cfg config.Config) (service *Service, err error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		url.QueryEscape(cfg.PostgresUser),
		url.QueryEscape(cfg.PostgresPassword),
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Exec("CREATE EXTENSION IF NOT EXISTS vector;").Error
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Content{}, &Chunk{})
	if err != nil {
		return nil, err
	}

	service = &Service{
		db: db,
	}
	return
}

func (s *Service) Close() {
	sqlDB, err := s.db.DB()
	if err != nil {
		log.Println("error closing database connection: ", err)
		return
	}

	if err = sqlDB.Close(); err != nil {
		log.Println("error closing database connection: ", err)
	}
}
