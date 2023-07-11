package handlers

import (
	errors_pkg "github/culinary_api/internal/delivery/http/errors"
	"github/culinary_api/internal/delivery/http/models"
	"github/culinary_api/internal/entity"
	"github/culinary_api/internal/usecase"

	"github.com/gin-gonic/gin"
)

type cookingStepHandlers struct {
	cookingStepsUseCase usecase.CookingSteps
	culinaryAggregator  usecase.CulinaryAggregator
}

func NewCookingStepHandlers(e *gin.RouterGroup, culinaryAggregator usecase.CulinaryAggregator, cookingStepsUseCase usecase.CookingSteps, middleware ...gin.HandlerFunc) {
	h := cookingStepHandlers{
		culinaryAggregator:  culinaryAggregator,
		cookingStepsUseCase: cookingStepsUseCase,
	}

	cookingStep := e.Group("/cooking-step", middleware...)
	{
		cookingStep.POST("/", h.CreateStep)
		cookingStep.PUT("/", h.UpdateStep)
		cookingStep.DELETE("/:id", h.DeleteStep)
	}
}

// @Security ApiKeyAuth
// @Router /v1/cooking-step [POST]
// @Summary Create cooking step
// @Description Create cooking step
// @Tags cooking-step
// @Accept json
// @Produce json
// @Param body body models.CreateStepRequest true "data"
// @Success 200 {object} models.CreateStepResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *cookingStepHandlers) CreateStep(c *gin.Context) {
	ctx := c.Request.Context()

	reqBody := &models.CreateStepRequest{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	i := &entity.CookingSteps{
		RecipeId:    reqBody.RecipeId,
		OrderNumber: reqBody.OrderNumber,
		Description: reqBody.Description,
		CookingTime: reqBody.CookingTime,
		Image:       reqBody.Image,
	}

	if err := r.culinaryAggregator.CreateCookingStep(ctx, i); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	c.JSON(200, &models.CreateStepResponse{
		Guid: i.Guid,
	})
}

// @Security ApiKeyAuth
// @Router /v1/cooking-step [PUT]
// @Summary Update cooking step
// @Description Update cooking step
// @Tags cooking-step
// @Accept json
// @Produce json
// @Param body body models.UpdateStepRequest true "data"
// @Success 200 {object} models.UpdateStepResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *cookingStepHandlers) UpdateStep(c *gin.Context) {
	ctx := c.Request.Context()

	reqBody := &models.UpdateStepRequest{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	i := &entity.CookingSteps{
		Guid:        reqBody.Guid,
		RecipeId:    reqBody.RecipeId,
		OrderNumber: reqBody.OrderNumber,
		Description: reqBody.Description,
		CookingTime: reqBody.CookingTime,
		Image:       reqBody.Image,
	}

	if err := r.culinaryAggregator.UpdateCookingStep(ctx, i); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	c.JSON(200, &models.UpdateStepResponse{
		Success: true,
	})
}

// @Security ApiKeyAuth
// @Router /v1/cooking-step/{id} [DELETE]
// @Summary Delete cooking step by id
// @Description Delete cooking step by id
// @Tags cooking-step
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.DeleteStepResponse
// @Failure 400 {object} models.ErrorBadRequest
func (r *cookingStepHandlers) DeleteStep(c *gin.Context) {
	ctx := c.Request.Context()

	if err := r.culinaryAggregator.DeleteCookingStep(ctx, c.Param("id")); err != nil {
		errors_pkg.Error(c, err)
		return
	}

	c.JSON(200, &models.DeleteStepResponse{
		Success: true,
	})

}
