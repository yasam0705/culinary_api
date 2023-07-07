package postgresql

import (
	"context"
	"github/culinary_api/pkg/postgres"

	"github.com/jackc/pgx/v4"
)

type base struct {
	db *postgres.DB
}

func NewBaseRepo(db *postgres.DB) *base {
	return &base{
		db: db,
	}
}

func (b *base) BeginTx(ctx context.Context) (context.Context, error) {
	return b.db.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})
}

func (b *base) Rollback(ctx context.Context) error {
	return b.db.Rollback(ctx)
}

func (b *base) Commit(ctx context.Context) error {
	return b.db.Commit(ctx)
}
