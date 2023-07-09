package models

type CreateStepRequest struct {
	RecipeId    string  `json:"recipe_id"`
	OrderNumber int     `json:"order_number"`
	Description string  `json:"description"`
	CookingTime float32 `json:"cooking_time"`
}

type CreateStepResponse struct {
	Guid string `json:"guid"`
}

type UpdateStepRequest struct {
	Guid        string  `json:"guid"`
	RecipeId    string  `json:"recipe_id"`
	OrderNumber int     `json:"order_number"`
	Description string  `json:"description"`
	CookingTime float32 `json:"cooking_time"`
}

type UpdateStepResponse struct {
	Success bool `json:"success"`
}

type DeleteStepResponse struct {
	Success bool `json:"success"`
}
