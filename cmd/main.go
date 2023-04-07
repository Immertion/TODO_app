package main

import (
	"log"

	"github.com/Immertion/TODO_app"
)

func main() {
	srv := new(TODO_app.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while ruunning http server: %s", err.Error())
	}
}
