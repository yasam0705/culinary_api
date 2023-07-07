package errors

import "net/http"

type customerError struct {
	Message string `json:"message"`
	ErrCode int    `json:"error_code"`
}

func (c *customerError) Error() string {
	return c.Message
}

func Error(err error) (int, *customerError) {
	return http.StatusBadRequest, &customerError{
		Message: err.Error(),
		ErrCode: http.StatusBadRequest,
	}
}
