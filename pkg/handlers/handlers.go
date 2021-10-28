package handlers

import (
	service "Skipper/pkg/servises"
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
	router.Use(corsMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/user-sign-up", h.signUp)
		auth.POST("/mentor-sign-up", h.mentorSignUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh-token", h.refreshToken)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.Group("/user")
		{
			api.GET("/user-data", h.GetUserData)
			api.POST("/user-mentor-sign-up", h.userToMentorSignUp)
			api.GET("/profile-picture/:filename", h.getUserProfilePicture)
		}
	}
	return router
}
