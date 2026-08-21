package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Service struct {
	db *bun.DB
}

func New(dsn string) (service *Service, closer func() error) {
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqlDb, pgdialect.New())
	closer = func() error {
		return sqlDb.Close()
	}

	return &Service{db: db}, closer
}
