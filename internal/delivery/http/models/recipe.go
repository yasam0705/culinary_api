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
