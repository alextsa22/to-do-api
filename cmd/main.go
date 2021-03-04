package main

import (
	"github.com/alextsa22/to-do-api/pkg/handler"
	"github.com/alextsa22/to-do-api/pkg/repository"
	"github.com/alextsa22/to-do-api/pkg/server"
	"github.com/alextsa22/to-do-api/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while  running http server: %s", err)
	}
}
