package models

type CreateStepRequest struct {
	RecipeId    string  `json:"recipe_id" validate:"required"`
	OrderNumber int     `json:"order_number" validate:"required"`
	Description string  `json:"description" validate:"required"`
	CookingTime float32 `json:"cooking_time"`
	Image       string  `json:"image"`
}

type CreateStepResponse struct {
	Guid string `json:"guid"`
}

type UpdateStepRequest struct {
	Guid        string  `json:"guid" validate:"required"`
	RecipeId    string  `json:"recipe_id" validate:"required"`
	OrderNumber int     `json:"order_number"`
	Description string  `json:"description"`
	CookingTime float32 `json:"cooking_time"`
	Image       string  `json:"image"`
}

type UpdateStepResponse struct {
	Success bool `json:"success"`
}

type DeleteStepResponse struct {
	Success bool `json:"success"`
}
