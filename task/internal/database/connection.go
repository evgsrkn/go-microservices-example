package database

import (
	"context"
	"fmt"
	"task/internal/config"

	"github.com/jackc/pgx/v5"
)

func New(cfg *config.Cfg) *pgx.Conn {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	return conn
}
