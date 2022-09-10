package handler

import (
	"github.com/TelitsynNikita/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api")
	{
		deals := api.Group("/deals")
		{
			deals.GET("/", h.getAllNewDeals)
			deals.GET("/:id", h.getOneDeal)
			deals.POST("/", h.createOneDeal)
			deals.DELETE("", h.deleteOneDeal)
			deals.PATCH("/", h.updateOneDeal)
			deals.PATCH("/change-deals", h.updateSomeDeals)
		}
	}

	return router
}
