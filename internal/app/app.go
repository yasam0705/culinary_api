package app

import (
	"github/culinary_api/config"
	http_router "github/culinary_api/internal/delivery/http"
	"github/culinary_api/internal/repository/postgresql"
	"github/culinary_api/internal/usecase"
	"github/culinary_api/pkg/logger"
	"github/culinary_api/pkg/postgres"
)

type App struct {
	cfg *config.Config
	db  *postgres.DB
	log logger.Logger
}

func NewApp(cfg *config.Config) (*App, error) {
	db, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}
	log, err := logger.New(cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		cfg: cfg,
		db:  db,
		log: log,
	}, nil
}

func (a *App) Run() error {

	// repo init
	baseRepo := postgresql.NewBaseRepo(a.db)
	recipeRepo := postgresql.NewRecipeRepo(a.db)
	ingredientsRepo := postgresql.NewIngredientsRepo(a.db)
	cookingStepsRepo := postgresql.NewCookingStepsRepo(a.db)
	recipeIngredientRepo := postgresql.NewRecipeIngredientRepo(a.db)
	culinaryAggregatorRepo := postgresql.NewCulinaryAggregatorRepo(a.db)

	// usecase init
	baseUseCase := usecase.NewBase(baseRepo)
	recipeUseCase := usecase.NewRecipe(recipeRepo)
	ingredientsUseCase := usecase.NewIngredients(ingredientsRepo)
	cookingStepsUseCase := usecase.NewCookingSteps(cookingStepsRepo)
	recipeIngredientUseCase := usecase.NewRecipeIngredient(recipeIngredientRepo)
	culinaryAggregatorUseCase := usecase.NewCulinaryAggregator(baseUseCase, culinaryAggregatorRepo, recipeUseCase, cookingStepsUseCase, ingredientsUseCase, recipeIngredientUseCase)

	services := http_router.CreateServices(
		cookingStepsUseCase,
		ingredientsUseCase,
		recipeUseCase,
		recipeIngredientUseCase,
		culinaryAggregatorUseCase,
	)

	// http init
	router, err := http_router.NewRouter(a.log, services)
	if err != nil {
		return err
	}

	return router.Run(":" + a.cfg.HttpPort)
}

func (a *App) Stop() error {
	a.db.Close()
	a.log.Close()
	return nil
}
