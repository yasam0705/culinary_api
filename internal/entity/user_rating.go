package entity

import "time"

type UserRating struct {
	Guid      string
	UserID    string
	RecipeID  string
	Rating    int8
	CreatedAt time.Time
}
