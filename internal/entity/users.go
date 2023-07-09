package entity

import "time"

type User struct {
	Guid      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
