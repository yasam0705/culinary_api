package models

type CulinaryAggregator struct {
	Recipe       *Recipe         `json:"recipe"`
	Ingredients  []*Ingredients  `json:"ingredients"`
	CookingSteps []*CookingSteps `json:"cooking_steps"`
}

type Recipe struct {
	Guid        string  `json:"guid"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	CookingTime float32 `json:"cooking_time"`
}

type Ingredients struct {
	Guid      string  `json:"guid"`
	Name      string  `json:"name"`
	Dimension string  `json:"dimension"`
	Count     float64 `json:"count"`
}

type CookingSteps struct {
	Guid        string  `json:"guid"`
	RecipeId    string  `json:"recipe_id"`
	OrderNumber int     `json:"order_number"`
	Description string  `json:"description"`
	CookingTime float32 `json:"cooking_time"`
}
