package models

// Guid
// Name
// Dimension
// Count
// RecipeId

type Ingridient struct {
	Guid      string `json:"guid"`
	Name      string `json:"name"`
	Dimension string `json:"dimension"`
}

type CreateIngredientRequest struct {
	Name      string  `json:"name"`
	Dimension string  `json:"dimension"`
	Count     float64 `json:"count"`
	RecipeId  string  `json:"recipe_id"`
}

type CreateIngredientResponse struct {
	Guid string `json:"guid"`
}

type UpdateIngredientRequest struct {
	Guid      string  `json:"guid"`
	Name      string  `json:"name"`
	Dimension string  `json:"dimension"`
	Count     float64 `json:"count"`
	RecipeId  string  `json:"recipe_id"`
}

type UpdateIngredientResponse struct {
	Success bool `json:"success"`
}

type DeleteIngredientResponse struct {
	Success bool `json:"success"`
}
