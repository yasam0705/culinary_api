package postgres

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
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
