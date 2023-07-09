package handlers

import (
	v1 "github/culinary_api/internal/delivery/http/handlers/v1"
	"github/culinary_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CookingSteps() usecase.CookingSteps
	Ingredients() usecase.Ingredients
	Recipe() usecase.Recipe
	RecipeIngredient() usecase.RecipeIngredient
	CulinaryAggregator() usecase.CulinaryAggregator
}

func NewHandlersV1(e *gin.Engine, srv Service) {
	v1Router := e.Group("/v1")

	v1.NewRecipeHandlers(v1Router, srv.Recipe(), srv.CulinaryAggregator())
	v1.NewIngridientHandlers(v1Router, srv.CulinaryAggregator(), srv.Ingredients())
	v1.NewCookingStepHandlers(v1Router, srv.CulinaryAggregator(), srv.CookingSteps())
}
