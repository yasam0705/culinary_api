package postgres

import (
	"context"
	"github/culinary_api/internal/entity"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

type txType struct{}

type DB struct {
	pool    *pgxpool.Pool
	Builder *SqlBuilder
}

func (db *DB) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	tx, ok := ctx.Value(txType{}).(pgx.Tx)
	if ok {
		return tx.Query(ctx, sql, args...)
	}
	return db.pool.Query(ctx, sql, args...)
}

func (db *DB) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	tx, ok := ctx.Value(txType{}).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, sql, args...)
	}
	return db.pool.QueryRow(ctx, sql, args...)
}

func (db *DB) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	tx, ok := ctx.Value(txType{}).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, sql, arguments...)
	}
	return db.pool.Exec(ctx, sql, arguments...)
}

func (db *DB) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (context.Context, error) {
	tx, err := db.pool.BeginTx(ctx, txOptions)
	return context.WithValue(ctx, txType{}, tx), err
}

func (db *DB) Rollback(ctx context.Context) error {
	tx := ctx.Value(txType{}).(pgx.Tx)
	return tx.Rollback(ctx)
}

func (db *DB) Commit(ctx context.Context) error {
	tx := ctx.Value(txType{}).(pgx.Tx)
	return tx.Commit(ctx)
}

func (db *DB) Close() {
	db.pool.Close()
}

func (db *DB) PgErr(err error) error {
	if err == nil {
		return nil
	}

	if pgErr, ok := err.(*pq.Error); ok {
		switch pgErr.Code {
		case "23505":
			return entity.ErrorAlreadyExists
		}
	}

	if err == pgx.ErrNoRows {
		return entity.ErrorNotFound
	}

	return err
}
