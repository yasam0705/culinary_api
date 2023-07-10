package models

import "errors"

var (
	NotAuthorized = errors.New("client not authorized")
)

type ErrorBadRequest struct {
	Message string `json:"message"`
	ErrCode int    `json:"error_code"`
}
