package entity

import (
	"time"
)

type Recipe struct {
	Guid        string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
