package main

import (
	"log"

	"github.com/Josh2604/go-notes-project/config"
	"github.com/Josh2604/go-notes-project/server"
)

func main() {
	/**
	Loads config file from config/config.yml
	*/
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	if err := server.Run("8080"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
