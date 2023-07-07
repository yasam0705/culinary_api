package postgres

import (
	"context"
	"fmt"
	"github/culinary_api/config"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func New(cfg *config.Config) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	pool, err := pgxpool.Connect(ctx, connString(cfg))
	if err != nil {
		return nil, err
	}

	sqlBuilder := newSqlBuilder()
	return &DB{
		pool:    pool,
		Builder: sqlBuilder,
	}, nil
}

func connString(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)
}
