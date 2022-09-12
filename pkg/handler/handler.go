package handler

import (
	"github.com/TelitsynNikita/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.userIdentity)
	{
		deals := api.Group("/deals")
		{
			deals.GET("/", h.getAllNewDeals)
			deals.POST("/get-own", h.getOwnDeals)
			deals.GET("/:id", h.getOneDeal)
			deals.POST("/", h.createOneDeal)
			deals.DELETE("/:id", h.deleteDeal)
			deals.PATCH("/", h.updateDealStatus)
			deals.PATCH("/update-bookkeeper-id", h.updateBookkeeperId)
		}
	}

	return router
}
