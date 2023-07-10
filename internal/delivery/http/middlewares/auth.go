package middlewares

import (
	errors_pkg "github/culinary_api/internal/delivery/http/errors"
	"github/culinary_api/internal/delivery/http/helper"
	"github/culinary_api/internal/delivery/http/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthM(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		token := c.GetHeader("Authorization")

		if err := helper.VerifyToken(m.cfg.Secret, token); err != nil {
			c.JSON(errors_pkg.Error(models.NotAuthorized))
			c.Abort()
			return
		}
	}
	c.Next()
}
