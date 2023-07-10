package errors

import (
	"github/culinary_api/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerError struct {
	Message string `json:"message"`
	ErrCode int    `json:"error_code"`
}

func (c *customerError) Error() string {
	return c.Message
}

func Error(c *gin.Context, err error) {
	var status int
	switch err {
	case entity.ErrorNotFound:
		status = http.StatusNotFound
	default:
		status = http.StatusBadRequest
	}
	c.JSON(http.StatusBadRequest, &customerError{
		Message: err.Error(),
		ErrCode: status,
	})
	c.Abort()
}
