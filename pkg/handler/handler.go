package handler

import (
	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	client pb.AuthServiceClient
}

func NewHandler(client pb.AuthServiceClient) *Handler {
	return &Handler{
		client: client,
	}
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

	account := router.Group("/account", h.userIdentity)
	{
		account.POST("/extend", h.extend)
		account.GET("/details", h.details)
	}

	return router
}
