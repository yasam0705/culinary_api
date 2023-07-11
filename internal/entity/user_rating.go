package entity

import (
	"errors"
	"time"
)

var (
	UserAlreadyVoted = errors.New("user has already voted")
)

type UserRating struct {
	Guid      string
	UserID    string
	RecipeID  string
	Rating    int8
	CreatedAt time.Time
}
