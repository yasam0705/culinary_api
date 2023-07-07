package http

import (
	"github/culinary_api/internal/delivery/http/handlers"
	"github/culinary_api/pkg/logger"

	_ "github/culinary_api/internal/delivery/http/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type router struct {
	log      logger.Logger
	Engine   *gin.Engine
	Services *services
}

func NewRouter(log logger.Logger, s *services) (*router, error) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	handlers.NewHandlersV1(r, s)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return &router{
		log:      log,
		Engine:   r,
		Services: s,
	}, nil
}

func (r *router) Run(port string) error {
	return r.Engine.Run(port)
}
