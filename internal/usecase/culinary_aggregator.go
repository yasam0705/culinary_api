package usecase

import (
	"context"
	"github/culinary_api/internal/entity"
	"log"
)

type CulinaryAggregator interface {
	CreateRecipe(ctx context.Context, m *entity.CulinaryAggregator) (err error)
	GetRecipe(ctx context.Context, filter map[string]string) (*entity.CulinaryAggregator, error)
	DeleteRecipe(ctx context.Context, filter map[string]string) (err error)
}

type CulinaryAggregatorRepo interface {
	Ingridients(ctx context.Context, filters map[string]string) ([]*entity.Ingredients, error)
}

type culinaryAggregator struct {
	*base
	repo             CulinaryAggregatorRepo
	recipe           Recipe
	cookingSteps     CookingSteps
	ingredients      Ingredients
	recipeIngredient RecipeIngredient
}

func NewCulinaryAggregator(
	base *base,
	repo CulinaryAggregatorRepo,
	recipe Recipe,
	cookingSteps CookingSteps,
	ingredients Ingredients,
	recipeIngredient RecipeIngredient,
) *culinaryAggregator {
	return &culinaryAggregator{
		base:             base,
		repo:             repo,
		recipe:           recipe,
		cookingSteps:     cookingSteps,
		ingredients:      ingredients,
		recipeIngredient: recipeIngredient,
	}
}

func (c *culinaryAggregator) CreateRecipe(ctx context.Context, m *entity.CulinaryAggregator) (err error) {
	contextTx, err := c.base.beginTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			log.Println("ROLLBACK TRANSACTION", c.base.rollback(contextTx))
		}
	}()

	if err = c.recipe.Create(contextTx, m.Recipe); err != nil {
		return err
	}

	for _, v := range m.CookingSteps {
		v.RecipeId = m.Recipe.Guid
	}

	if err = c.cookingSteps.BatchCreate(contextTx, m.CookingSteps); err != nil {
		return err
	}

	if err = c.ingredients.BatchCreate(contextTx, m.Ingredients); err != nil {
		return err
	}

	var recipeIngredient = make([]*entity.RecipeIngredient, 0, len(m.Ingredients))
	for _, v := range m.Ingredients {
		recipeIngredient = append(recipeIngredient, &entity.RecipeIngredient{
			RecipeId:     m.Recipe.Guid,
			IngredientId: v.Guid,
			Count:        v.Count,
		})
	}

	if err = c.recipeIngredient.BatchCreate(contextTx, recipeIngredient); err != nil {
		return err
	}

	return c.commit(contextTx)
}

func (c *culinaryAggregator) GetRecipe(ctx context.Context, filter map[string]string) (*entity.CulinaryAggregator, error) {
	recipe, err := c.recipe.Get(ctx, filter)
	if err != nil {
		return nil, err
	}

	steps, err := c.cookingSteps.List(ctx, 0, 0, filter)
	if err != nil {
		return nil, err
	}

	ingredients, err := c.repo.Ingridients(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &entity.CulinaryAggregator{
		Recipe:       recipe,
		CookingSteps: steps,
		Ingredients:  ingredients,
	}, nil
}

func (c *culinaryAggregator) DeleteRecipe(ctx context.Context, filter map[string]string) (err error) {
	// TRANSACTION
	contextTx, err := c.base.beginTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			log.Println("ROLLBACK TRANSACTION", c.base.rollback(contextTx))
		}
	}()

	if err = c.recipeIngredient.Delete(contextTx, filter); err != nil {
		return err
	}

	if err = c.cookingSteps.Delete(contextTx, filter); err != nil {
		return err
	}

	if err = c.recipe.Delete(contextTx, filter); err != nil {
		return err
	}

	return c.commit(contextTx)
}

func (c *culinaryAggregator) UpdateRecipe(ctx context.Context, m *entity.CulinaryAggregator) (err error) {
	// TRANSACTION
	contextTx, err := c.base.beginTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			log.Println("ROLLBACK TRANSACTION", c.base.rollback(contextTx))
		}
	}()

	if err = c.recipe.Update(contextTx, m.Recipe); err != nil {
		return err
	}

	return c.commit(contextTx)
}
