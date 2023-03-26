package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/venomuz/alif-task/internal/config"
	"github.com/venomuz/alif-task/internal/service"
)

type Handler struct {
	services *service.Services
	cfg      *config.Config
}

func NewHandler(services *service.Services, cfg *config.Config) *Handler {
	return &Handler{
		services: services,
		cfg:      cfg,
	}
}
func (h *Handler) Init(api *gin.RouterGroup) {

	v1 := api.Group("/v1")
	{
		h.initSettingsRoutes(v1)
	}

}
