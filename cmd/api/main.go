package main

import (
	"log"

	"github.com/Josh2604/go-notes-project/server"
)

func main() {
	if err := server.Run("8080"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
