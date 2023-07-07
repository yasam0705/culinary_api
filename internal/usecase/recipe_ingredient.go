package usecase

import (
	"context"
	"github/culinary_api/internal/entity"
)

type RecipeIngredient interface {
	Create(ctx context.Context, m *entity.RecipeIngredient) error
	Update(ctx context.Context, m *entity.RecipeIngredient) error
	Get(ctx context.Context, m map[string]string) (*entity.RecipeIngredient, error)
	List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.RecipeIngredient, error)
	BatchCreate(ctx context.Context, m []*entity.RecipeIngredient) error
	Delete(ctx context.Context, filter map[string]string) error
}

type RecipeIngredientRepo interface {
	Create(ctx context.Context, m *entity.RecipeIngredient) error
	Update(ctx context.Context, m *entity.RecipeIngredient) error
	FindOne(ctx context.Context, m map[string]string) (*entity.RecipeIngredient, error)
	FindAll(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.RecipeIngredient, error)
	Delete(ctx context.Context, filter map[string]string) error
}

type recipeIngredient struct {
	repo RecipeIngredientRepo
}

func NewRecipeIngredient(repo RecipeIngredientRepo) *recipeIngredient {
	return &recipeIngredient{
		repo: repo,
	}
}

func (r *recipeIngredient) Create(ctx context.Context, m *entity.RecipeIngredient) error {
	return r.repo.Create(ctx, m)
}

func (r *recipeIngredient) Get(ctx context.Context, m map[string]string) (*entity.RecipeIngredient, error) {
	return r.repo.FindOne(ctx, m)
}

func (r *recipeIngredient) List(ctx context.Context, limit, offset uint64, m map[string]string) ([]*entity.RecipeIngredient, error) {
	return r.repo.FindAll(ctx, limit, offset, m)
}

func (r *recipeIngredient) Update(ctx context.Context, m *entity.RecipeIngredient) error {
	return r.repo.Update(ctx, m)
}

func (r *recipeIngredient) BatchCreate(ctx context.Context, m []*entity.RecipeIngredient) error {
	for _, v := range m {
		if err := r.repo.Create(ctx, v); err != nil {
			return err
		}
	}
	return nil
}

func (r *recipeIngredient) Delete(ctx context.Context, filter map[string]string) error {
	return r.repo.Delete(ctx, filter)
}
