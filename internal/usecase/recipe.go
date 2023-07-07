package usecase

import (
	"context"
	"github/culinary_api/internal/entity"
	"time"

	"github.com/google/uuid"
)

type Recipe interface {
	Create(ctx context.Context, m *entity.Recipe) error
	Update(ctx context.Context, m *entity.Recipe) error
	Get(ctx context.Context, m map[string]string) (*entity.Recipe, error)
	List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.Recipe, error)
	Delete(ctx context.Context, filter map[string]string) error
}

type RecipeRepo interface {
	Create(ctx context.Context, m *entity.Recipe) error
	Update(ctx context.Context, m *entity.Recipe) error
	FindOne(ctx context.Context, m map[string]string) (*entity.Recipe, error)
	FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.Recipe, error)
	Delete(ctx context.Context, filter map[string]string) error
}

type recipe struct {
	repo RecipeRepo
}

func NewRecipe(repo RecipeRepo) *recipe {
	return &recipe{
		repo: repo,
	}
}

func (r *recipe) Create(ctx context.Context, m *entity.Recipe) error {
	r.beforeCreate(m)
	return r.repo.Create(ctx, m)
}

func (r *recipe) Get(ctx context.Context, m map[string]string) (*entity.Recipe, error) {
	return r.repo.FindOne(ctx, m)
}

func (r *recipe) List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.Recipe, error) {
	return r.repo.FindAll(ctx, limit, offset, m)
}

func (r *recipe) Update(ctx context.Context, m *entity.Recipe) error {
	r.beforeUpdate(m)
	return r.repo.Update(ctx, m)
}

func (r *recipe) beforeUpdate(m *entity.Recipe) {
	m.UpdatedAt = time.Now()
}

func (r *recipe) beforeCreate(m *entity.Recipe) {
	m.Guid = uuid.New().String()
	m.CreatedAt = time.Now()
	m.UpdatedAt = m.CreatedAt
}

func (r *recipe) Delete(ctx context.Context, filter map[string]string) error {
	return r.repo.Delete(ctx, filter)
}
