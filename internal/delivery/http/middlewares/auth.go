package middlewares

import (
	errors_pkg "github/culinary_api/internal/delivery/http/errors"
	"github/culinary_api/internal/delivery/http/helper"
	"github/culinary_api/internal/delivery/http/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *middleware) AuthM(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		token := c.GetHeader("Authorization")

		sub, err := helper.VerifyToken(m.cfg.Secret, token)
		if err != nil {
			log.Println(err)
			errors_pkg.Error(c, models.NotAuthorized)
			return
		}
		c.Set("user_id", sub)
	}
	c.Next()
}
