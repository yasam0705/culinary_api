package http

import "github/culinary_api/internal/usecase"

type services struct {
	cookingSteps       usecase.CookingSteps
	ingredients        usecase.Ingredients
	recipe             usecase.Recipe
	recipeIngredient   usecase.RecipeIngredient
	culinaryAggregator usecase.CulinaryAggregator
	auth               usecase.Auth
}

func CreateServices(cookingSteps usecase.CookingSteps, ingredients usecase.Ingredients, recipe usecase.Recipe, recipeIngredient usecase.RecipeIngredient, culinaryAggregator usecase.CulinaryAggregator, auth usecase.Auth) *services {
	return &services{
		cookingSteps:       cookingSteps,
		ingredients:        ingredients,
		recipe:             recipe,
		recipeIngredient:   recipeIngredient,
		culinaryAggregator: culinaryAggregator,
		auth:               auth,
	}
}

func (s *services) CookingSteps() usecase.CookingSteps {
	return s.cookingSteps
}

func (s *services) Ingredients() usecase.Ingredients {
	return s.ingredients
}

func (s *services) Recipe() usecase.Recipe {
	return s.recipe
}

func (s *services) RecipeIngredient() usecase.RecipeIngredient {
	return s.recipeIngredient
}

func (s *services) CulinaryAggregator() usecase.CulinaryAggregator {
	return s.culinaryAggregator
}

func (s *services) Auth() usecase.Auth {
	return s.auth
}
