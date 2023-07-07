package models

type ErrorBadRequest struct {
	Message string `json:"message"`
	ErrCode int    `json:"error_code"`
}
