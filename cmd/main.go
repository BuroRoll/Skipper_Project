package main

import (
	"Skipper/pkg/handlers"
	"Skipper/pkg/models"
	"Skipper/pkg/repository"
	service "Skipper/pkg/servises"
)

func main() {
	//srv := new(backend.Server)
	db := models.GetDB()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlerses := handlers.NewHandler(services)

	if err := handlers.InitSocket(); err != nil {
		panic(err)
	}
	handlerses.SocketEvents()
	go handlers.SocketServer.Serve()
	defer handlers.SocketServer.Close()
	handlerses.InitRoutes()
}
