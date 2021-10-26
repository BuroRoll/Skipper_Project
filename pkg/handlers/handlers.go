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

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware())

	auth := router.Group("/auth")
	{

		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh-token", h.refreshToken)
		auth.GET("/logout", h.logout)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/status", h.GetStatus)
	}

	router.GET("/ping", func(g *gin.Context) {
		g.JSON(200, gin.H{"text": "Hello from public"})
	})

	//router.GET("/secured/status", h.userIdentity, h.GetStatus)
	return router
}
