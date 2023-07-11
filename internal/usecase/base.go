package usecase

import "context"

type BaseRepo interface {
	BeginTx(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

//go:generate mockgen -destination=tests/mocks/base.go -package=mocks -source=base.go
type Base interface {
	BeginTx(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

type base struct {
	repo BaseRepo
}

func NewBase(repo BaseRepo) *base {
	return &base{repo: repo}
}

func (b *base) BeginTx(ctx context.Context) (context.Context, error) {
	return b.repo.BeginTx(ctx)
}

func (b *base) Rollback(ctx context.Context) error {
	return b.repo.Rollback(ctx)
}

func (b *base) Commit(ctx context.Context) error {
	return b.repo.Commit(ctx)
}
