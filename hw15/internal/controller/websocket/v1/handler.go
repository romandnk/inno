package v1

import (
	"github.com/gin-gonic/gin"

	"chat/internal/service"
)

type Handler struct {
	engine   *gin.Engine
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	h.engine = router

	// api := router.Group("/api/v1")
	//{
	//	// urls management group
	//	//urls := api.Group("/urls")
	//	//{
	//	//	//urlroute.NewUrlRoutes(urls, h.services.URL)
	//	//}
	//}

	return h.engine
}
