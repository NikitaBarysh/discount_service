package handler

import (
	"github.com/NikitaBarysh/discount_service.git/internal/middleware"
	"github.com/NikitaBarysh/discount_service.git/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services   *service.Service
	middleware middleware.Middleware
}

func NewHandler(newService *service.Service) *Handler {
	return &Handler{services: newService}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/api/user")
	{
		auth.POST("/register", h.signUp)
		auth.POST("/login", h.signIn)
	}

	user := router.Group("/api/user", h.middleware.UserIdentity)
	{
		user.POST("/orders", h.setOrder)
		user.GET("/orders", h.getOrders)
		user.GET("/balance", h.getBalance)
		user.POST("/balance/withdraw", h.useWithdraw)
		user.GET("withdrawals", h.getWithdraw)
	}

	return router
}
