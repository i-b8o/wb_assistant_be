package handler

import (
	"github.com/bogach-ivan/wb_assistant_be/be/api/pkg/service"
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

	// Create groups
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/confirmation/:token", h.confirmation)
		auth.POST("/resend", h.resend)
		auth.POST("/set", h.set)
		auth.POST("/password-reset", h.passwordReset)
	}

	return router
}
