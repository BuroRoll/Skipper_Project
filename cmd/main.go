package main

import (
	backend "Skipper"
	"Skipper/pkg/handlers"
	"Skipper/pkg/models"
	"Skipper/pkg/repository"
	service "Skipper/pkg/servises"
	"log"
)

func main() {
	srv := new(backend.Server)
	db := models.GetDB()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlerses := handlers.NewHandler(services)
	if err := srv.Run("8000", handlerses.InitRoutes()); err != nil {
		log.Fatalf("Error run server: %s", err)
	}
}
