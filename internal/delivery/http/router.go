package http

import (
	"github/culinary_api/config"
	"github/culinary_api/internal/delivery/http/handlers"
	"github/culinary_api/internal/delivery/http/middlewares"
	"github/culinary_api/pkg/logger"

	_ "github/culinary_api/internal/delivery/http/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type router struct {
	Log      logger.Logger
	Cfg      *config.Config
	engine   *gin.Engine
	services *services
}

// NewRoute
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRouter(cfg *config.Config, log logger.Logger, s *services) (*router, error) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	m := middlewares.NewMiddleware(cfg)

	handlers.NewHandlersV1(r, cfg, s, m)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return &router{
		Log:      log,
		Cfg:      cfg,
		engine:   r,
		services: s,
	}, nil
}

func (r *router) Run(port string) error {
	return r.engine.Run(port)
}
