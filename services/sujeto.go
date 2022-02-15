package services

import (
	"fmt"
	"time"

	"github.com/LNMMusic/fce/models"
)

func CreateSujeto() *models.Sujeto {
	return &models.Sujeto {
		Nombre:			"Lucas",
		Edad:			23,
		Personeria:		0,
		Nacionalidad:	0,

		Actividades:	CreateActividades(5),
	}
}

func CreateActividades(quanty int) []*models.Actividad {
	var actividades []*models.Actividad

	for i:=0; i<quanty; i++ {
		actividades = append(actividades, &models.Actividad {
			Accion:			fmt.Sprintf("Accion [%d]", i),
			Objeto:			fmt.Sprintf("Objecto [%d]", i),
			Descripcion:	fmt.Sprintf("Descripcion [%d]", i),

			Inicio:			time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
			Fin:			time.Now(),

			Ubicacion:		fmt.Sprintf("Ubicacion [%d]", i),
		})
	}

	return actividades
}