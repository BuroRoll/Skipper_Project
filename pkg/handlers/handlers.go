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
			user.POST("/user-mentor-sign-up", h.userToMentorSignUp)
		}
		catalog := api.Group("/catalog")
		{
			catalog.POST("/create-catalog", h.createCatalog)
		}
	}
	router.GET("/user/profile-picture/:filename", h.getUserProfilePicture)

	catalog := router.Group("/catalog")
	{
		catalog.GET("/", h.getCatalog)
		catalog.GET("/main-section", h.mainSection)
	}
	return router
}
