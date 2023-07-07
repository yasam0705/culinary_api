package entity

type CulinaryAggregator struct {
	Recipe       *Recipe
	Ingredients  []*Ingredients
	CookingSteps []*CookingSteps
}
