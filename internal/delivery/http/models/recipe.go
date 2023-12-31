package models

type CreateAggregatorRequest struct {
	Recipe       *CreateRecipeRequest         `json:"recipe"`
	Ingredients  []*CreateIngredientsRequest  `json:"ingredients"`
	CookingSteps []*CreateCookingStepsRequest `json:"cooking_steps"`
}

type CreateRecipeRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateIngredientsRequest struct {
	Name      string  `json:"name"`
	Dimension string  `json:"dimension"`
	Count     float64 `json:"count"`
}

type CreateCookingStepsRequest struct {
	OrderNumber int    `json:"order_number"`
	Description string `json:"description"`
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
}

type IngredientsUpdate struct {
	Guid      string  `json:"guid"`
	Name      string  `json:"name"`
	Dimension string  `json:"dimension"`
	Count     float64 `json:"count"`
}

type CookingStepsUpdate struct {
	Guid        string `json:"guid"`
	RecipeId    string `json:"recipe_id"`
	OrderNumber int    `json:"order_number"`
	Description string `json:"description"`
}

type UpdateRecipeRequest struct {
	Recipe       *RecipeUpdate         `json:"recipe"`
	Ingredients  []*IngredientsUpdate  `json:"ingredients"`
	CookingSteps []*CookingStepsUpdate `json:"cooking_steps"`
}

type UpdateRecipeResponse struct {
	Success bool `json:"succress"`
}
