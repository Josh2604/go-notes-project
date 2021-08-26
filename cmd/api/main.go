package main

import (
	"log"

	"github.com/Josh2604/go-notes-project/config"
	"github.com/Josh2604/go-notes-project/server"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	if err := server.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
