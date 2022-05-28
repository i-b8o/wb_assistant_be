package handler

import (
	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/bogach-ivan/wb_assistant_be/docs"
)

type Handler struct {
	authClient pb.AuthServiceClient
	mailClient pb.MailServiceClient
}

func NewHandler(authClient pb.AuthServiceClient, mailClient pb.MailServiceClient) *Handler {
	return &Handler{
		authClient: authClient,
		mailClient: mailClient,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Create groups
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/confirmation/:token", h.confirmation)
		auth.POST("/resend", h.resend)
		auth.POST("/set", h.set)
		auth.POST("/recover", h.passwordRecover)
	}

	account := router.Group("/account", h.userIdentity)
	{
		account.POST("/update", h.update)
		account.POST("/update-email-verification-token", h.updateEmailVerificationToken)
		account.GET("/details", h.details)
	}

	return router
}
