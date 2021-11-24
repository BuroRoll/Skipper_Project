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

	api := router.Group("/api")
	{
		user := api.Group("/user", h.userIdentity)
		{
			user.GET("/user-data", h.GetUserData)

			user.POST("/user-mentor-sign-up", h.userToMentorSignUp)
			user.GET("/user-verify-email", h.UserVerifyEmail)
			user.POST("/update-base-profile-data", h.UpdateBaseProfileData)
			user.POST("/update-profile-picture", h.UpdateProfilePicture)

			communication := api.Group("/communication")
			{
				communication.GET("/user-communications", h.GetUserCommunications)
				communication.GET("/messenger-list", h.GetMessengers)
				communication.POST("/create-user-communication", h.CreateUserCommunication)
			}
			education := api.Group("/education")
			{
				education.GET("/user-education", h.GetUserEducations)
				education.POST("/add-user-education", h.AddUserEducation)
			}
			workExperience := api.Group("/work-experience")
			{
				workExperience.POST("/add-user-work-experience", h.AddUserWorkExperience)
				workExperience.GET("/user-work-experience", h.GetUserWorkExperience)
			}
		}
		publicUser := api.Group("/public-user")
		{
			publicUser.GET("/mentor/:id", h.GetMentorData)
			publicUser.GET("/menti/:id", h.GetMentiData)
		}
		catalog := api.Group("/catalog")
		{
			catalog.GET("/", h.getCatalog)
			catalog.GET("/main-section", h.mainSection)
		}
		api.GET("/user/profile-picture/:filename", h.GetUserProfilePicture)
		api.GET("/verify-email", h.verifyEmail)
	}

	return router
}
