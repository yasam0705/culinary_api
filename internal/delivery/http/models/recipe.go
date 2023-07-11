package models

type CreateAggregatorRequest struct {
	Recipe       *CreateRecipeRequest         `json:"recipe"`
	Ingredients  []*CreateIngredientsRequest  `json:"ingredients"`
	CookingSteps []*CreateCookingStepsRequest `json:"cooking_steps"`
}

type CreateRecipeRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image       string `json:"image"`
}

type CreateIngredientsRequest struct {
	Name      string  `json:"name" validate:"required"`
	Dimension string  `json:"dimension" validate:"required"`
	Count     float64 `json:"count" validate:"required"`
}

type CreateCookingStepsRequest struct {
	OrderNumber int     `json:"order_number" validate:"required"`
	Description string  `json:"description" validate:"required"`
	CookingTime float32 `json:"cooking_time"`
	Image       string  `json:"image"`
}

type CreateAggregatorResponse struct {
	Guid string `json:"guid"`
}

type DeleteRecipeResponse struct {
	Success bool `json:"succress"`
}

type RecipeUpdate struct {
	Guid        string `json:"guid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type IngredientsUpdate struct {
	Guid      string  `json:"guid"`
	Name      string  `json:"name"`
	Dimension string  `json:"dimension"`
	Count     float64 `json:"count"`
}

type CookingStepsUpdate struct {
	Guid        string  `json:"guid"`
	RecipeId    string  `json:"recipe_id"`
	OrderNumber int     `json:"order_number"`
	Description string  `json:"description"`
	CookingTime float32 `json:"cooking_time"`
	Image       string  `json:"image"`
}

type UpdateRecipeRequest struct {
	Recipe       *RecipeUpdate         `json:"recipe"`
	Ingredients  []*IngredientsUpdate  `json:"ingredients"`
	CookingSteps []*CookingStepsUpdate `json:"cooking_steps"`
}

type UpdateRecipeResponse struct {
	Success bool `json:"succress" `
}

type RecipeRatingRequest struct {
	Rating   int8   `json:"rating" minimum:"0" maximum:"6"`
	RecipeId string `json:"recipe_id"`
}

type RecipeRatingResponse struct {
	Success bool `json:"succress"`
}
