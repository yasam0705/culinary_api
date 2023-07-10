package entity

import (
	"time"
)

type Recipe struct {
	Guid            string
	Title           string
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CookingTime     float32
	Rating          int64
	NumberOfRatings int64
	OverallRating   float64
}
