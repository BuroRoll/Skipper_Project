package handlers

import (
	service "Skipper/pkg/servises"
	//"github.com/auth0/go-jwt-middleware"
	//"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

//import (
//	service "Skipper/pkg/servises"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/gin-gonic/gin"
//)

//
//import (
//	service "Skipper/pkg/servises"
//	"github.com/gin-gonic/gin"
//	"github.com/itsjamie/gin-cors"
//	"time"
//)
//
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

//
//
//func (h *Handler) InitRoutes() *gin.Engine {
//	router := gin.New()
//
//	router.Use(cors.Middleware(cors.Config{
//		Origins:        "*",
//		Methods:        "GET, PUT, POST, DELETE",
//		RequestHeaders: "Origin, Authorization, Content-Type",
//		ExposedHeaders: "",
//		MaxAge: 50 * time.Second,
//		Credentials: true,
//		ValidateHeaders: true,
//	}))
//	router.Use(cors.Middleware(cors.Config{
//		Origins:        "*",
//		Methods:        "*",
//		RequestHeaders: "*",
//		ExposedHeaders: "*",
//		MaxAge: 50 * time.Second,
//	}))
//
//
//	api := router.Group("/api")
//	{
//		api.Group("/auth")
//		{
//			api.POST("/sign-up", h.signUp)
//			api.POST("/sign-in", h.signIn)
//			api.GET("/logout", h.logout)
//		}
//		api.Group("/status")
//		{
//			api.GET("/status", h.GetStatus)
//		}
//	}
//
//
//
//	//api := router.Group("/api", h.userIdentity)
//	//{
//	//
//	//}
//	return router
//}

//var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
//	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
//		fmt.Println(token.Claims)
//		return true, nil
//	},
//	SigningMethod: jwt.SigningMethodHS256,
//})
//
//func checkJWT() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		jwtMid := *jwtMiddleware
//		if err := jwtMid.CheckJWT(c.Writer, c.Request); err != nil {
//			c.AbortWithStatus(401)
//		}
//	}
//}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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
	r := gin.Default()
	r.Use(corsMiddleware())

	auth := r.Group("/auth")
	{

		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/logout", h.logout)

	}

	r.GET("/ping", func(g *gin.Context) {
		g.JSON(200, gin.H{"text": "Hello from public"})
	})

	r.GET("/secured/ping", h.userIdentity, func(g *gin.Context) {
		g.JSON(200, gin.H{"text": "Hello from private"})
	})

	return r
}
