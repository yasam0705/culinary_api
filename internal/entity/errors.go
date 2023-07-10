package entity

import "errors"

var (
	ErrorNotFound      = errors.New("no rows returned")
	ErrorAlreadyExists = errors.New("row already exists")
)
