package main

import (
	"fmt"

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
	var persona = &services.Sujeto {
		Nombre:		"Lucas",
		Edad:		23,
		Personeria:	3,
	}
	
	fmt.Printf("Hi %s, you have %d and you are %s",
		persona.Nombre,
		persona.Edad,
		persona.Personeria.String(),
	)
}