package usecase

import (
	"context"
	"github/culinary_api/internal/entity"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=tests/mocks/cooking_steps.go -package=mocks -source=cooking_steps.go
type CookingSteps interface {
	Create(ctx context.Context, m *entity.CookingSteps) error
	Update(ctx context.Context, m *entity.CookingSteps) error
	Get(ctx context.Context, m map[string]string) (*entity.CookingSteps, error)
	List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.CookingSteps, error)
	BatchCreate(ctx context.Context, m []*entity.CookingSteps) error
	Delete(ctx context.Context, filter map[string]string) error
	BatchUpdate(ctx context.Context, m []*entity.CookingSteps) error
}

type CookingStepsRepo interface {
	Create(ctx context.Context, m *entity.CookingSteps) error
	Update(ctx context.Context, m *entity.CookingSteps) error
	FindOne(ctx context.Context, m map[string]string) (*entity.CookingSteps, error)
	FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.CookingSteps, error)
	Delete(ctx context.Context, filter map[string]string) error
}

type cookingSteps struct {
	repo      CookingStepsRepo
	batchSize int
}

func NewCookingSteps(repo CookingStepsRepo) *cookingSteps {
	return &cookingSteps{
		repo:      repo,
		batchSize: 50,
	}
}

func (i *cookingSteps) Create(ctx context.Context, m *entity.CookingSteps) error {
	i.beforeCreate(m)
	return i.repo.Create(ctx, m)
}

func (i *cookingSteps) Get(ctx context.Context, m map[string]string) (*entity.CookingSteps, error) {
	return i.repo.FindOne(ctx, m)
}

func (i *cookingSteps) List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.CookingSteps, error) {
	return i.repo.FindAll(ctx, limit, offset, m)
}

func (i *cookingSteps) Update(ctx context.Context, m *entity.CookingSteps) error {
	return i.repo.Update(ctx, m)
}

func (i *cookingSteps) beforeCreate(m *entity.CookingSteps) {
	m.Guid = uuid.New().String()
}

func (i *cookingSteps) BatchCreate(ctx context.Context, m []*entity.CookingSteps) error {
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

func (i *cookingSteps) BatchUpdate(ctx context.Context, m []*entity.CookingSteps) error {
	// refactor
	if len(m) == 0 {
		return nil
	}
	for _, v := range m {
		if err := i.repo.Update(ctx, v); err != nil {
			return err
		}
	}
	return nil
}

func (i *cookingSteps) Delete(ctx context.Context, filter map[string]string) error {
	return i.repo.Delete(ctx, filter)
}
