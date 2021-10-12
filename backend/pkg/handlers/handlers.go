package handlers

import (
	service "Skipper/pkg/servises"
	"github.com/gin-gonic/contrib/sessions"
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

	router.Use(sessions.Sessions("session", sessions.NewCookieStore([]byte("secret"))))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/logout", h.logout)
	}

	api := router.Group("/api")
	{
		status := api.Group("/status")
		{
			status.GET("/", h.GetStatus)
		}
	}
	return router
}
