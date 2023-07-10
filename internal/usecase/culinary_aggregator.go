package usecase

import (
	"context"
	"errors"
	"fmt"
	"github/culinary_api/internal/entity"
	math_pkg "github/culinary_api/pkg/math"
	"log"
)

type CulinaryAggregator interface {
	CreateRecipe(ctx context.Context, m *entity.CulinaryAggregator) (err error)
	GetRecipe(ctx context.Context, filter map[string]string) (*entity.CulinaryAggregator, error)
	DeleteRecipe(ctx context.Context, filter map[string]string) (err error)
	UpdateRecipe(ctx context.Context, m *entity.CulinaryAggregator) (err error)

	// ingridient
	CreateIngredient(ctx context.Context, recipeId string, m *entity.Ingredients) (err error)
	UpdateIngredient(ctx context.Context, recipeId string, m *entity.Ingredients) (err error)
	DeleteIngredient(ctx context.Context, recipeId, ingridientId string) error

	// cooking-step
	CreateCookingStep(ctx context.Context, m *entity.CookingSteps) (err error)
	UpdateCookingStep(ctx context.Context, m *entity.CookingSteps) (err error)
	DeleteCookingStep(ctx context.Context, id string) (err error)

	// rating
	AddRating(ctx context.Context, userId, recipeId string, rating int8) (err error)
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
	userRatings      UserRatings
}

func NewCulinaryAggregator(
	base *base,
	repo CulinaryAggregatorRepo,
	recipe Recipe,
	cookingSteps CookingSteps,
	ingredients Ingredients,
	recipeIngredient RecipeIngredient,
	userRatings UserRatings,
) *culinaryAggregator {
	return &culinaryAggregator{
		base:             base,
		repo:             repo,
		recipe:           recipe,
		cookingSteps:     cookingSteps,
		ingredients:      ingredients,
		recipeIngredient: recipeIngredient,
		userRatings:      userRatings,
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

	recipeIngridients := make([]*entity.RecipeIngredient, 0, len(m.Ingredients))
	for _, v := range m.Ingredients {
		recipeIngridients = append(recipeIngridients, &entity.RecipeIngredient{
			RecipeId:     m.Recipe.Guid,
			IngredientId: v.Guid,
			Count:        v.Count,
		})
	}

	if c.recipeIngredient.BatchUpdate(ctx, recipeIngridients); err != nil {
		return err
	}

	if err = c.cookingSteps.BatchUpdate(ctx, m.CookingSteps); err != nil {
		return err
	}

	return c.commit(contextTx)
}

// INGRIDIENT
func (c *culinaryAggregator) CreateIngredient(ctx context.Context, recipeId string, m *entity.Ingredients) (err error) {
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

	if err = c.ingredients.Create(contextTx, m); err != nil {
		return err
	}

	l := &entity.RecipeIngredient{
		RecipeId:     recipeId,
		IngredientId: m.Guid,
		Count:        m.Count,
	}

	if err = c.recipeIngredient.Create(contextTx, l); err != nil {
		return err
	}
	return c.commit(contextTx)
}

func (c *culinaryAggregator) UpdateIngredient(ctx context.Context, recipeId string, m *entity.Ingredients) (err error) {
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

	if err = c.ingredients.Update(contextTx, m); err != nil {
		return err
	}

	l := &entity.RecipeIngredient{
		RecipeId:     recipeId,
		IngredientId: m.Guid,
		Count:        m.Count,
	}

	if err = c.recipeIngredient.Update(contextTx, l); err != nil {
		return err
	}
	return c.commit(contextTx)
}

func (c *culinaryAggregator) DeleteIngredient(ctx context.Context, recipeId, ingridientId string) error {
	return c.recipeIngredient.Delete(ctx, map[string]string{
		"recipe_id":     recipeId,
		"ingredient_id": ingridientId,
	})
}

// COOKING STEP
func (c *culinaryAggregator) CreateCookingStep(ctx context.Context, m *entity.CookingSteps) (err error) {
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

	if err = c.cookingSteps.Create(contextTx, m); err != nil {
		return err
	}

	recipe, err := c.recipe.Get(contextTx, map[string]string{
		"recipe_id": m.RecipeId,
	})
	if err != nil {
		return err
	}
	recipe.CookingTime += m.CookingTime

	if err = c.recipe.Update(contextTx, recipe); err != nil {
		return err
	}

	return c.commit(contextTx)
}

func (c *culinaryAggregator) UpdateCookingStep(ctx context.Context, m *entity.CookingSteps) (err error) {
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

	oldStep, err := c.cookingSteps.Get(contextTx, map[string]string{
		"guid": m.Guid,
	})
	if err != nil {
		return err
	}

	if err = c.cookingSteps.Update(contextTx, m); err != nil {
		return err
	}

	recipe, err := c.recipe.Get(contextTx, map[string]string{
		"recipe_id": m.RecipeId,
	})
	if err != nil {
		return err
	}
	recipe.CookingTime += (m.CookingTime - oldStep.CookingTime)

	if err = c.recipe.Update(contextTx, recipe); err != nil {
		return err
	}

	return c.commit(contextTx)
}

func (c *culinaryAggregator) DeleteCookingStep(ctx context.Context, id string) (err error) {
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

	filter := map[string]string{
		"guid": id,
	}

	step, err := c.cookingSteps.Get(contextTx, filter)
	if err != nil {
		return err
	}

	recipe, err := c.recipe.Get(contextTx, map[string]string{
		"recipe_id": step.RecipeId,
	})
	if err != nil {
		return err
	}
	recipe.CookingTime -= step.CookingTime

	if err = c.recipe.Update(contextTx, recipe); err != nil {
		return err
	}

	if err = c.cookingSteps.Delete(contextTx, filter); err != nil {
		return err
	}

	return c.commit(contextTx)
}

// RATING
func (c *culinaryAggregator) AddRating(ctx context.Context, userId, recipeId string, rating int8) (err error) {
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

	userRating, err := c.userRatings.Get(contextTx, map[string]string{
		"user_id": userId,
	})
	if err != nil && !errors.Is(err, entity.ErrorNotFound) {
		return err
	}
	if userRating != nil {
		return fmt.Errorf("user has already voted")
	}

	userRating = &entity.UserRating{
		UserID:   userId,
		RecipeID: recipeId,
		Rating:   int8(rating),
	}

	if err = c.userRatings.Create(contextTx, userRating); err != nil {
		return err
	}

	recipe, err := c.recipe.Get(contextTx, map[string]string{
		"recipe_id": recipeId,
	})
	if err != nil {
		return err
	}

	recipe.Rating += int64(rating)
	recipe.NumberOfRatings += 1
	recipe.OverallRating = math_pkg.MathRound(float64(recipe.Rating) / float64(recipe.NumberOfRatings))

	if err = c.recipe.Update(contextTx, recipe); err != nil {
		return err
	}

	return c.commit(contextTx)
}
