package main

import (
	"github.com/alextsa22/to-do-api/pkg/server"
	"log"
)

func main() {
	srv := new(server.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while  running http server: %s", err)
	}
}
