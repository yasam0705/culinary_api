package handlers

import (
	"github/culinary_api/config"
	v1 "github/culinary_api/internal/delivery/http/handlers/v1"
	"github/culinary_api/internal/delivery/http/middlewares"
	"github/culinary_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CookingSteps() usecase.CookingSteps
	Ingredients() usecase.Ingredients
	Recipe() usecase.Recipe
	RecipeIngredient() usecase.RecipeIngredient
	CulinaryAggregator() usecase.CulinaryAggregator
	Auth() usecase.Auth
}

func NewHandlersV1(e *gin.Engine, cfg *config.Config, srv Service, middlewares middlewares.Middleware) {
	v1Router := e.Group("/v1")

	v1.NewRecipeHandlers(v1Router, srv.Recipe(), srv.CulinaryAggregator(), middlewares.AuthM)
	v1.NewIngridientHandlers(v1Router, srv.CulinaryAggregator(), srv.Ingredients(), middlewares.AuthM)
	v1.NewCookingStepHandlers(v1Router, srv.CulinaryAggregator(), srv.CookingSteps(), middlewares.AuthM)
	v1.NewAuthHandlers(v1Router, cfg.Secret, cfg.AcessTokenTTL, cfg.RefreshTokenTTL, srv.Auth())
}
