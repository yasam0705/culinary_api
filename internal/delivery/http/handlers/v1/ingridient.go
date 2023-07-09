package handlers

import (
	errors_pkg "github/culinary_api/internal/delivery/http/errors"
	"github/culinary_api/internal/delivery/http/helper"
	"github/culinary_api/internal/delivery/http/models"

	"github/culinary_api/internal/entity"
	"github/culinary_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ingridientHandlers struct {
	ingredientsUseCase usecase.Ingredients
	culinaryAggregator usecase.CulinaryAggregator
}

func NewIngridientHandlers(e *gin.RouterGroup, culinaryAggregator usecase.CulinaryAggregator, ingredientsUseCase usecase.Ingredients) {
	h := ingridientHandlers{
		ingredientsUseCase: ingredientsUseCase,
		culinaryAggregator: culinaryAggregator,
	}

	ingridient := e.Group("/recipe-ingridient")
	{
		ingridient.GET("/", h.IngredientList)
		ingridient.POST("/", h.CreateIngredient)
		ingridient.PUT("/", h.UpdateIngredient)
		ingridient.DELETE("/", h.DeleteIngredient)
	}
}

// @Router /v1/recipe-ingridient [GET]
// @Summary ingridient list
// @Description ingridient list
// @Tags recipe-ingridient
// @Accept json
// @Produce json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {array} models.Ingridient
// @Failure 400 {object} models.ErrorBadRequest
func (r *ingridientHandlers) IngredientList(c *gin.Context) {
	ctx := c.Request.Context()

	params, err := helper.GetQueryParams(c)
	if err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	list, err := r.ingredientsUseCase.List(ctx, params.GetLimit(), params.GetOffset(), params.GetFilter())
	if err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	response := make([]*models.Ingridient, 0, len(list))
	for _, v := range list {
		response = append(response, &models.Ingridient{
			Guid:      v.Guid,
			Name:      v.Name,
			Dimension: v.Dimension,
		})
	}
	c.JSON(200, response)
}

// @Router /v1/recipe-ingridient [POST]
// @Summary Create ingridient
// @Description Create ingridient
// @Tags recipe-ingridient
// @Accept json
// @Produce json
// @Param body body models.CreateIngredientRequest true "data"
// @Success 200 {object} models.CreateIngredientResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *ingridientHandlers) CreateIngredient(c *gin.Context) {
	ctx := c.Request.Context()
	reqBody := &models.CreateIngredientRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}
	i := &entity.Ingredients{
		Name:      reqBody.Name,
		Dimension: reqBody.Dimension,
		Count:     reqBody.Count,
	}

	if err := r.culinaryAggregator.CreateIngredient(ctx, reqBody.RecipeId, i); err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	c.JSON(200, &models.CreateIngredientResponse{
		Guid: i.Guid,
	})
}

// @Router /v1/recipe-ingridient [PUT]
// @Summary Update ingridient
// @Description Update ingridient
// @Tags recipe-ingridient
// @Accept json
// @Produce json
// @Param body body models.UpdateIngredientRequest true "data"
// @Success 200 {object} models.UpdateIngredientResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *ingridientHandlers) UpdateIngredient(c *gin.Context) {
	ctx := c.Request.Context()
	reqBody := &models.UpdateIngredientRequest{}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	i := &entity.Ingredients{
		Guid:      reqBody.Guid,
		Name:      reqBody.Name,
		Dimension: reqBody.Dimension,
		Count:     reqBody.Count,
	}

	if err := r.culinaryAggregator.UpdateIngredient(ctx, reqBody.RecipeId, i); err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	c.JSON(200, &models.UpdateIngredientResponse{
		Success: true,
	})
}

// @Router /v1/recipe-ingridient [DELETE]
// @Summary Delete ingridient by id
// @Description Delete ingridient by id
// @Tags recipe-ingridient
// @Accept json
// @Produce json
// @Param recipe_id query string true "recipe_id"
// @Param ingredient_id query string true "ingredient_id"
// @Success 200 {array} models.DeleteIngredientResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *ingridientHandlers) DeleteIngredient(c *gin.Context) {
	ctx := c.Request.Context()
	params, err := helper.GetQueryParams(c)
	if err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	err = r.culinaryAggregator.DeleteRecipe(ctx, params.GetFilter())
	if err != nil {
		c.JSON(errors_pkg.Error(err))
		return
	}

	c.JSON(200, &models.DeleteIngredientResponse{
		Success: true,
	})
}
