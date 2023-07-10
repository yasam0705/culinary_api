package usecase

import (
	"context"
	"github/culinary_api/internal/entity"
	"time"

	"github.com/google/uuid"
)

type User interface {
	Create(ctx context.Context, m *entity.User) error
	Update(ctx context.Context, m *entity.User) error
	Get(ctx context.Context, m map[string]string) (*entity.User, error)
	List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.User, error)
	Delete(ctx context.Context, filter map[string]string) error
}

type UserRepo interface {
	Create(ctx context.Context, m *entity.User) error
	Update(ctx context.Context, m *entity.User) error
	FindOne(ctx context.Context, m map[string]string) (*entity.User, error)
	FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.User, error)
	Delete(ctx context.Context, filter map[string]string) error
}

type user struct {
	repo UserRepo
}

func NewUser(repo UserRepo) *user {
	return &user{
		repo: repo,
	}
}

func (u *user) beforeCreate(m *entity.User) {
	m.Guid = uuid.New().String()
	m.CreatedAt = time.Now()
	m.UpdatedAt = m.CreatedAt
}

func (u *user) beforeUpdate(m *entity.User) {
	m.UpdatedAt = time.Now()
}

func (u *user) Create(ctx context.Context, m *entity.User) error {
	u.beforeCreate(m)
	return u.repo.Create(ctx, m)
}
func (u *user) Update(ctx context.Context, m *entity.User) error {
	u.beforeUpdate(m)
	return u.repo.Update(ctx, m)
}
func (u *user) Get(ctx context.Context, m map[string]string) (*entity.User, error) {
	return u.repo.FindOne(ctx, m)
}
func (u *user) List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.User, error) {
	return u.repo.FindAll(ctx, limit, offset, m)
}
func (u *user) Delete(ctx context.Context, filter map[string]string) error {
	return u.repo.Delete(ctx, filter)
}
