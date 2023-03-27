package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/venomuz/alif-task/docs"
	"github.com/venomuz/alif-task/internal/config"
	"github.com/venomuz/alif-task/internal/service"
	v1 "github.com/venomuz/alif-task/internal/transport/rest/v1"
	"github.com/venomuz/alif-task/pkg/middleware"
	"net/http"
)

type Handler struct {
	services *service.Services
	cfg      config.Config
}

func NewHandler(services *service.Services, cfg config.Config) *Handler {
	return &Handler{
		services: services,
		cfg:      cfg,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./docs/swagger.yaml",
		SpecPath:    "/redoc/swagger.yaml",
		DocsPath:    "/redoc",
	}

	router.Use(
		ginredoc.New(doc),
		gin.Recovery(),
		gin.Logger(),
		middleware.New(GinCorsMiddleware()),
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", h.cfg.HTTP.Host, h.cfg.HTTP.Port)

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.cfg)

	{
		api := router.Group("/api")
		{
			handlerV1.Init(api)
		}
	}

}
