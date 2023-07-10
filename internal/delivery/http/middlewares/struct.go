package middlewares

import (
	"github/culinary_api/config"

	"github.com/gin-gonic/gin"
)

type middleware struct {
	cfg *config.Config
}

type Middleware interface {
	AuthM(c *gin.Context)
}

func NewMiddleware(cfg *config.Config) *middleware {
	return &middleware{
		cfg: cfg,
	}
}
