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
			user.POST("/user-verify-email", h.UserVerifyEmail)
			user.POST("/update-base-profile-data", h.UpdateBaseProfileData)
			user.POST("/update-profile-picture", h.UpdateProfilePicture)
			user.POST("/update-specialization", h.UpdateSpecialization)

			communication := api.Group("/communication")
			{
				communication.DELETE("user-communication/:id", h.DeleteUserCommunication)
				communication.GET("/user-communications", h.GetUserCommunications)
				communication.GET("/messenger-list", h.GetMessengers)
				communication.POST("/create-user-communication", h.CreateUserCommunication)
			}
			education := api.Group("/education")
			{
				education.DELETE("/user-education/:id", h.DeleteUserEducation)
				education.GET("/user-education", h.GetUserEducations)
				education.POST("/add-user-education", h.AddUserEducation)
			}
			workExperience := api.Group("/work-experience")
			{
				workExperience.DELETE("/user-work-experience/:id", h.DeleteUserWorkExperience)
				workExperience.POST("/add-user-work-experience", h.AddUserWorkExperience)
				workExperience.GET("/user-work-experience", h.GetUserWorkExperience)
			}
			otherInfo := api.Group("/other-info")
			{
				otherInfo.DELETE("/user-other-info/:id", h.DeleteUserOtherInfo)
				otherInfo.POST("/add-user-other-info", h.AddUserOtherInfo)
				otherInfo.GET("/user-other-info", h.GetUserOtherInfo)
			}
			class := api.Group("/class")
			{
				class.POST("/class", h.CreateUserClass)
				class.POST("/theoretic-class", h.CreateTheoreticClass)
				class.POST("/practic-class", h.CreatePracticClass)
				class.POST("/key-class", h.CreateKeyClass)

				class.DELETE("/class/:id", h.DeleteClass)
				class.DELETE("/theoretic-class/:id", h.DeleteTheoreticClass)
				class.DELETE("/practic-class/:id", h.DeletePracticClass)
				class.DELETE("/key-class/:id", h.DeleteKeyClass)

				class.PUT("/class", h.UpdateClass)
				class.PUT("/theoretic-class", h.UpdateTheoreticClass)
				class.PUT("/practic-class", h.UpdatePracticClass)
				class.PUT("/key-class", h.UpdateKeyClass)

				class.GET("/user-classes", h.GetUserClasses)

				booking := class.Group("/booking")
				{
					booking.GET("", h.GetClassData)
					booking.POST("", h.BookClass)
				}
			}
		}

	}
	publicApi := router.Group("/public-api")
	{
		publicUser := publicApi.Group("/public-user")
		{
			publicUser.GET("/mentor/:id", h.GetMentorData)
			publicUser.GET("/menti/:id", h.GetMentiData)
		}
		catalog := publicApi.Group("/catalog")
		{
			catalog.GET("/", h.GetCatalog)
			catalog.GET("/main-section", h.GetMainSection)
			catalog.GET("/child", h.GetCatalogChild)
			catalog.GET("/classes", h.GetClasses)
		}
		publicApi.GET("/user/profile-picture/:filename", h.GetUserProfilePicture)
	}

	router.GET("/verify-email", h.verifyEmail)

	return router
}
