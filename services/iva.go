package services

import (
	"log"

	"github.com/LNMMusic/fce/models"
)

func SujetoIVAResume() {
	var sujeto = models.SujetoNew()
	
	models.IVA.Resume(sujeto)
	for _, operacion := range sujeto.Operaciones {
		log.Printf("operacion -> %#v\n", operacion.Taxes)
	}
}