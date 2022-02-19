package main

import (
	"github.com/LNMMusic/fce/config"
	"github.com/LNMMusic/fce/models"
	"github.com/LNMMusic/fce/services"
)

func main() {
	// Env
	config.EnvLoad()

	// DB
	// ...

	// Cache
	models.IVA.Load()

	// App
	services.SujetoIVAResume()
}