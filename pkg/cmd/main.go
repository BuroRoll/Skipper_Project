package main

import (
	"Skipper/pkg/handlers"
	"Skipper/pkg/models"
	"Skipper/pkg/repository"
	service "Skipper/pkg/servises"
)

// @title Skipper Backend
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db := models.GetDB()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlerses := handlers.NewHandler(services)
	handlers.InitSocket()
	handlerses.SocketEvents()
	go handlers.SocketServer.Serve()
	defer handlers.SocketServer.Close()
	handlerses.InitRoutes()
}
