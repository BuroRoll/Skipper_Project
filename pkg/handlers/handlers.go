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
		user := api.Group("/user")
		{
			user.GET("/user-data", h.GetUserData)
			user.GET("/user-communications", h.GetUserCommunications)
			user.POST("/user-mentor-sign-up", h.userToMentorSignUp)
			user.GET("/user-verify-email", h.UserVerifyEmail)
			user.POST("/update-base-profile-data", h.UpdateBaseProfileData)
		}
		communication := api.Group("/communication")
		{
			communication.GET("/messenger-list", h.GetMessengers)
			communication.POST("/create-user-communication", h.CreateUserCommunication)
		}
		catalog := api.Group("/catalog")
		{
			catalog.POST("/create-catalog", h.createCatalog)
		}
	}
	router.GET("/user/profile-picture/:filename", h.GetUserProfilePicture)
	router.GET("/verify-email", h.verifyEmail)

	catalog := router.Group("/catalog")
	{
		catalog.GET("/", h.getCatalog)
		catalog.GET("/main-section", h.mainSection)
	}
	return router
}
