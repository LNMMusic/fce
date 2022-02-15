package services

import (
	"time"
)

// Enums [const block -> tuple of types enumerated by index]
type EnumPersoneria int
const (
	Juridica	EnumPersoneria = iota
	Humana
	Sucesion
)
func (e EnumPersoneria) String() string {
	switch e {
	case Juridica:
		return "Juridica"
	case Humana:
		return "Humana"
	case Sucesion:
		return "Sucesion"
	default:
		return ""
	}
}


// Structs
type Actividad struct {
	Accion			string
	Objeto			string
	Descripcion		string

	Inicio			time.Time
	Fin				time.Time

	Ubicacion		string
}

type Sujeto struct {
	Nombre			string
	Edad			uint8
	Personeria		EnumPersoneria		// takes the const type based on index
	Nacionalidad	string

	// Aspectos
	Actividades		[]Actividad
}