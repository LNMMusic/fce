package main

import (
	"log"

	"github.com/LNMMusic/fce/config"
	"github.com/LNMMusic/fce/client"
	"github.com/LNMMusic/fce/services"
)

func main() {
	// Env
	config.EnvLoad()

	// DB
	client.Mongo.ConnectClient()
	defer client.Mongo.DisconnectClient()

	// App
	sujeto := services.CreateSujeto()
	log.Printf("data -> %#v", sujeto)
}