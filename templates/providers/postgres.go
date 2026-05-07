package providers

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type IPostgresConfig interface {
	GetUri() string
}

func NewPostgreSQL(config IPostgresConfig) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), config.GetUri())
	if err != nil {
		panic(err)
	}
	return conn
}
