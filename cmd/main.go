package main

import (
	"github.com/alextsa22/to-do-api/pkg/handler"
	"github.com/alextsa22/to-do-api/pkg/server"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while  running http server: %s", err)
	}
}
