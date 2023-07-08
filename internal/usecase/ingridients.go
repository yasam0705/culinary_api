package usecase

import (
	"context"
	"github/culinary_api/internal/entity"

	"github.com/google/uuid"
)

type Ingredients interface {
	Create(ctx context.Context, m *entity.Ingredients) error
	Update(ctx context.Context, m *entity.Ingredients) error
	Get(ctx context.Context, m map[string]string) (*entity.Ingredients, error)
	List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.Ingredients, error)
	BatchCreate(ctx context.Context, m []*entity.Ingredients) error
}

type IngredientsRepo interface {
	Create(ctx context.Context, m *entity.Ingredients) error
	Update(ctx context.Context, m *entity.Ingredients) error
	FindOne(ctx context.Context, m map[string]string) (*entity.Ingredients, error)
	FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.Ingredients, error)
}

type ingredients struct {
	repo      IngredientsRepo
	batchSize int
}

func NewIngredients(repo IngredientsRepo) *ingredients {
	return &ingredients{
		repo:      repo,
		batchSize: 50,
	}
}

func (i *ingredients) Create(ctx context.Context, m *entity.Ingredients) error {
	i.beforeCreate(m)
	return i.repo.Create(ctx, m)
}

func (i *ingredients) Get(ctx context.Context, m map[string]string) (*entity.Ingredients, error) {
	return i.repo.FindOne(ctx, m)
}

func (i *ingredients) List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.Ingredients, error) {
	return i.repo.FindAll(ctx, limit, offset, m)
}

func (i *ingredients) Update(ctx context.Context, m *entity.Ingredients) error {
	return i.repo.Update(ctx, m)
}

func (i *ingredients) beforeCreate(m *entity.Ingredients) {
	m.Guid = uuid.New().String()
}

func (i *ingredients) BatchCreate(ctx context.Context, m []*entity.Ingredients) error {
	// refactor
	if len(m) == 0 {
		return nil
	}
	for _, v := range m {
		i.beforeCreate(v)
		if err := i.repo.Create(ctx, v); err != nil {
			return err
		}

	}
	return nil
}
