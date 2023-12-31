package handlers

import (
	errors_pkg "github/culinary_api/internal/delivery/http/errors"
	"github/culinary_api/internal/delivery/http/handlers/pkg"
	"github/culinary_api/internal/delivery/http/models"

	"github/culinary_api/internal/entity"
	"github/culinary_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type recipeHandlers struct {
	recipeUsecase      usecase.Recipe
	culinaryAggregator usecase.CulinaryAggregator
}

func NewRecipeHandlers(e *gin.RouterGroup, recipeUsecase usecase.Recipe, culinaryAggregator usecase.CulinaryAggregator) {
	h := recipeHandlers{
		recipeUsecase:      recipeUsecase,
		culinaryAggregator: culinaryAggregator,
	}

	recipe := e.Group("/recipe")
	{
		recipe.GET("/", h.RecipeList)
		recipe.GET("/:id", h.Recipe)
		recipe.POST("/", h.CreateRecipe)
		recipe.PUT("/", h.UpdateRecipe)
		recipe.DELETE("/:id", h.DeleteRecipe)
	}
}

// @Router /v1/recipe [GET]
// @Summary Recipe List
// @Description Recipe List
// @Tags aggregator
// @Accept json
// @Produce json
// @Param limit query int true "limit"
// @Param offset query int true "offset"
// @Success 200 {array} models.Recipe
// @Failure 400 {object} models.ErrorBadRequest
func (r *recipeHandlers) RecipeList(c *gin.Context) {
	ctx := c.Request.Context()

	params, err := pkg.GetQueryParams(c)
	if err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	list, err := r.recipeUsecase.List(ctx, params.GetLimit(), params.GetOffset(), params.GetFilter())
	if err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	response := make([]*models.Recipe, 0, len(list))
	for _, v := range list {
		response = append(response, &models.Recipe{
			Guid:        v.Guid,
			Title:       v.Title,
			Description: v.Description,
			CreatedAt:   pkg.TimeToStrRFC3339(v.CreatedAt),
			UpdatedAt:   pkg.TimeToStrRFC3339(v.UpdatedAt),
		})
	}
	c.JSON(200, response)
}

// @Router /v1/recipe/{id} [GET]
// @Summary Recipe by id
// @Description Recipe by id
// @Tags aggregator
// @Accept json
// @Produce json
// @Param id path string true "recipe_id"
// @Success 200 {array} models.CulinaryAggregator
// @Failure 400 {object} models.ErrorBadRequest
func (r *recipeHandlers) Recipe(c *gin.Context) {
	ctx := c.Request.Context()

	result, err := r.culinaryAggregator.GetRecipe(ctx, map[string]string{
		"recipe_id": c.Param("id"),
	})
	if err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	ingredients := make([]*models.Ingredients, 0, len(result.Ingredients))
	for _, v := range result.Ingredients {
		ingredients = append(ingredients, &models.Ingredients{
			Guid:      v.Guid,
			Name:      v.Name,
			Dimension: v.Dimension,
			Count:     v.Count,
		})
	}
	steps := make([]*models.CookingSteps, 0, len(result.CookingSteps))
	for _, v := range result.CookingSteps {
		steps = append(steps, &models.CookingSteps{
			Guid:        v.Guid,
			RecipeId:    v.RecipeId,
			OrderNumber: v.OrderNumber,
			Description: v.Description,
		})
	}
	c.JSON(200, &models.CulinaryAggregator{
		Recipe: &models.Recipe{
			Guid:        result.Recipe.Guid,
			Title:       result.Recipe.Title,
			Description: result.Recipe.Description,
			CreatedAt:   pkg.TimeToStrRFC3339(result.Recipe.CreatedAt),
			UpdatedAt:   pkg.TimeToStrRFC3339(result.Recipe.UpdatedAt),
		},
		Ingredients:  ingredients,
		CookingSteps: steps,
	})
}

// @Router /v1/recipe [POST]
// @Summary Create recipe
// @Description Create recipe
// @Tags aggregator
// @Accept json
// @Produce json
// @Param body body models.CreateAggregatorRequest true "data"
// @Success 200 {object} models.CreateAggregatorResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *recipeHandlers) CreateRecipe(c *gin.Context) {
	ctx := c.Request.Context()
	reqBody := &models.CreateAggregatorRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}
	m := r.convertToEntityCreate(reqBody)

	if err := r.culinaryAggregator.CreateRecipe(ctx, m); err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	c.JSON(200, models.CreateAggregatorResponse{
		Guid: m.Recipe.Guid,
	})

}

// @Router /v1/recipe [PUT]
// @Summary Update recipe
// @Description Update recipe
// @Tags aggregator
// @Accept json
// @Produce json
// @Param body body models.UpdateRecipeRequest true "data"
// @Success 200 {object} models.UpdateRecipeResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *recipeHandlers) UpdateRecipe(c *gin.Context) {
	ctx := c.Request.Context()

	reqBody := &models.UpdateRecipeRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	m := r.convertToEntityUpdate(reqBody)

	if err := r.culinaryAggregator.UpdateRecipe(ctx, m); err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	c.JSON(200, &models.UpdateRecipeResponse{
		Success: true,
	})

}

// @Router /v1/recipe/{id} [DELETE]
// @Summary Delete recipe by id
// @Description Delete recipe by id
// @Tags aggregator
// @Accept json
// @Produce json
// @Param id path string true "recipe_id"
// @Success 200 {array} models.DeleteRecipeResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *recipeHandlers) DeleteRecipe(c *gin.Context) {
	ctx := c.Request.Context()

	err := r.culinaryAggregator.DeleteRecipe(ctx, map[string]string{
		"recipe_id": c.Param("id"),
	})
	if err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	c.JSON(200, &models.DeleteRecipeResponse{
		Success: true,
	})
}

func (r *recipeHandlers) convertToEntityCreate(reqBody *models.CreateAggregatorRequest) *entity.CulinaryAggregator {
	ingredients := make([]*entity.Ingredients, 0, len(reqBody.Ingredients))
	for _, v := range reqBody.Ingredients {
		ingredients = append(ingredients, &entity.Ingredients{
			Name:      v.Name,
			Dimension: v.Dimension,
			Count:     v.Count,
		})
	}

	steps := make([]*entity.CookingSteps, 0, len(reqBody.CookingSteps))
	for _, v := range reqBody.CookingSteps {
		steps = append(steps, &entity.CookingSteps{
			OrderNumber: v.OrderNumber,
			Description: v.Description,
		})
	}

	res := &entity.CulinaryAggregator{
		Recipe: &entity.Recipe{
			Title:       reqBody.Recipe.Title,
			Description: reqBody.Recipe.Description,
		},
		Ingredients:  ingredients,
		CookingSteps: steps,
	}

	return res
}

func (r *recipeHandlers) convertToEntityUpdate(reqBody *models.UpdateRecipeRequest) *entity.CulinaryAggregator {
	ingredients := make([]*entity.Ingredients, 0, len(reqBody.Ingredients))
	for _, v := range reqBody.Ingredients {
		ingredients = append(ingredients, &entity.Ingredients{
			Guid:      v.Guid,
			Name:      v.Name,
			Dimension: v.Dimension,
			Count:     v.Count,
		})
	}

	steps := make([]*entity.CookingSteps, 0, len(reqBody.CookingSteps))
	for _, v := range reqBody.CookingSteps {
		steps = append(steps, &entity.CookingSteps{
			Guid:        v.Guid,
			RecipeId:    v.RecipeId,
			OrderNumber: v.OrderNumber,
			Description: v.Description,
		})
	}

	res := &entity.CulinaryAggregator{
		Recipe: &entity.Recipe{
			Guid:        reqBody.Recipe.Guid,
			Title:       reqBody.Recipe.Title,
			Description: reqBody.Recipe.Description,
		},
		Ingredients:  ingredients,
		CookingSteps: steps,
	}

	return res
}
