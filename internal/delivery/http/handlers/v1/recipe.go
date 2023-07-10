package handlers

import (
	"fmt"
	errors_pkg "github/culinary_api/internal/delivery/http/errors"
	"github/culinary_api/internal/delivery/http/helper"
	"github/culinary_api/internal/delivery/http/models"

	"github/culinary_api/internal/entity"
	"github/culinary_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type recipeHandlers struct {
	recipeUsecase      usecase.Recipe
	culinaryAggregator usecase.CulinaryAggregator
}

func NewRecipeHandlers(e *gin.RouterGroup, recipeUsecase usecase.Recipe, culinaryAggregator usecase.CulinaryAggregator, middleware ...gin.HandlerFunc) {
	h := recipeHandlers{
		recipeUsecase:      recipeUsecase,
		culinaryAggregator: culinaryAggregator,
	}

	recipe := e.Group("/recipe", middleware...)
	{
		recipe.GET("/", h.RecipeList)
		recipe.GET("/:id", h.Recipe)
		recipe.POST("/", h.CreateRecipe)
		recipe.PUT("/", h.UpdateRecipe)
		recipe.DELETE("/:id", h.DeleteRecipe)
		recipe.POST("/rating", h.RecipeRating)
	}
}

// @Router /v1/recipe [GET]
// @Summary Recipe List
// @Description Recipe List
// @Tags aggregator
// @Accept json
// @Produce json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Param cooking_time_from query string false "cooking_time_from"
// @Param cooking_time_to query string false "cooking_time_to"
// @Param ingridients query string false "guid,guid"
// @Param rating_from query float64 false "rating_from"
// @Param rating_to query float64 false "rating_to"
// @Param order_rating query int false "order_rating"
// @Success 200 {array} models.Recipe
// @Failure 400 {object} models.ErrorBadRequest
func (r *recipeHandlers) RecipeList(c *gin.Context) {
	ctx := c.Request.Context()

	params, err := helper.GetQueryParams(c)
	if err != nil {
		errors_pkg.Error(c, err)
		return
	}

	list, err := r.recipeUsecase.List(ctx, params.GetLimit(), params.GetOffset(), params.GetFilter())
	if err != nil {
		errors_pkg.Error(c, err)
		return
	}

	response := make([]*models.Recipe, 0, len(list))
	for _, v := range list {
		response = append(response, &models.Recipe{
			Guid:            v.Guid,
			Title:           v.Title,
			Description:     v.Description,
			CreatedAt:       helper.TimeToStrRFC3339(v.CreatedAt),
			UpdatedAt:       helper.TimeToStrRFC3339(v.UpdatedAt),
			CookingTime:     v.CookingTime,
			Rating:          v.Rating,
			NumberOfRatings: v.NumberOfRatings,
			OverallRating:   v.OverallRating,
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
		errors_pkg.Error(c, err)
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
			CookingTime: v.CookingTime,
		})
	}
	c.JSON(200, &models.CulinaryAggregator{
		Recipe: &models.Recipe{
			Guid:            result.Recipe.Guid,
			Title:           result.Recipe.Title,
			Description:     result.Recipe.Description,
			CreatedAt:       helper.TimeToStrRFC3339(result.Recipe.CreatedAt),
			UpdatedAt:       helper.TimeToStrRFC3339(result.Recipe.UpdatedAt),
			CookingTime:     result.Recipe.CookingTime,
			Rating:          result.Recipe.Rating,
			NumberOfRatings: result.Recipe.NumberOfRatings,
			OverallRating:   result.Recipe.OverallRating,
		},
		Ingredients:  ingredients,
		CookingSteps: steps,
	})
}

// @Security ApiKeyAuth
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
		errors_pkg.Error(c, err)
		return
	}
	m := r.convertToEntityCreate(reqBody)

	if err := r.culinaryAggregator.CreateRecipe(ctx, m); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	c.JSON(200, models.CreateAggregatorResponse{
		Guid: m.Recipe.Guid,
	})
}

// @Security ApiKeyAuth
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
		errors_pkg.Error(c, err)
		return
	}
	m := r.convertToEntityUpdate(reqBody)

	if err := r.culinaryAggregator.UpdateRecipe(ctx, m); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	c.JSON(200, &models.UpdateRecipeResponse{
		Success: true,
	})
}

// @Security ApiKeyAuth
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
		errors_pkg.Error(c, err)
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

	var allCookingTime float32
	steps := make([]*entity.CookingSteps, 0, len(reqBody.CookingSteps))
	for _, v := range reqBody.CookingSteps {
		allCookingTime += v.CookingTime
		steps = append(steps, &entity.CookingSteps{
			OrderNumber: v.OrderNumber,
			Description: v.Description,
			CookingTime: v.CookingTime,
		})
	}

	res := &entity.CulinaryAggregator{
		Recipe: &entity.Recipe{
			Title:       reqBody.Recipe.Title,
			Description: reqBody.Recipe.Description,
			CookingTime: allCookingTime,
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

	var allCookingTime float32
	steps := make([]*entity.CookingSteps, 0, len(reqBody.CookingSteps))
	for _, v := range reqBody.CookingSteps {
		allCookingTime += v.CookingTime
		steps = append(steps, &entity.CookingSteps{
			Guid:        v.Guid,
			RecipeId:    v.RecipeId,
			OrderNumber: v.OrderNumber,
			Description: v.Description,
			CookingTime: v.CookingTime,
		})
	}

	res := &entity.CulinaryAggregator{
		Recipe: &entity.Recipe{
			Guid:        reqBody.Recipe.Guid,
			Title:       reqBody.Recipe.Title,
			Description: reqBody.Recipe.Description,
			CookingTime: allCookingTime,
		},
		Ingredients:  ingredients,
		CookingSteps: steps,
	}

	return res
}

// @Security ApiKeyAuth
// @Router /v1/recipe/rating [POST]
// @Summary Recipe rating
// @Description Recipe rating
// @Tags rating
// @Accept json
// @Produce json
// @Param body body models.RecipeRatingRequest true "data"
// @Success 200 {object} models.RecipeRatingResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *recipeHandlers) RecipeRating(c *gin.Context) {
	ctx := c.Request.Context()
	reqBody := &models.RecipeRatingRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	if reqBody.Rating <= 0 && reqBody.Rating > 5 {
		errors_pkg.Error(c, fmt.Errorf("rating must be from 1 to 5"))
		return
	}

	userId, ok := c.Get("user_id")
	if !ok {
		errors_pkg.Error(c, fmt.Errorf("user not found"))
		return
	}

	if err := r.culinaryAggregator.AddRating(ctx, userId.(string), reqBody.RecipeId, reqBody.Rating); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	c.JSON(200, models.RecipeRatingResponse{
		Success: true,
	})
}
