package main

import (
	"Skipper/pkg/handlers"
	"Skipper/pkg/models"
	"Skipper/pkg/repository"
	service "Skipper/pkg/servises"
)

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
