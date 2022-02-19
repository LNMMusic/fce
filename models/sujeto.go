package models

import (
	"fmt"
	"time"
)


// STRUCTS
type Tiempo struct {
	Inicio			time.Time			`bson:"inicio,omitempty"`
	Fin				time.Time			`bson:"fin,omitempty"`
}
func (t *Tiempo) Hours() float64 {
	return t.Fin.Sub(t.Inicio).Hours()
}

type Tax struct {
	IVA		float64
	IIGG	float64
}

type Operacion struct {
	Tipo			EnumOperacion		`bson:"tipo,omitempty"`
	Objeto			EnumObjeto			`bson:"objeto,omitempty"`
	Titulo			EnumTitulo			`bson:"titulo,omitempty"`
	Descripcion		string				`bson:"descripcion,omitempty"`

	Ubicacion		string				`bson:"ubicacion,omitempty"`
	Duracion		*Tiempo				`bson:"duracion,omitempty"`
	Residencia		bool				`bson:"residencia,omitempty"`

	Taxes			Tax					`bson:"taxes,omitempty"`
}
// constructor [init | gen]
func OperacionNew(id int) *Operacion {
	return &Operacion {
		Tipo:			0,
		Objeto:			0,
		Titulo:			0,
		Descripcion:	fmt.Sprintf("Descripcion [%d]", id),

		Ubicacion:		fmt.Sprintf("Ubicacion [%d]", id),
		Duracion:		&Tiempo {
			Inicio:		time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
			Fin:		time.Now(),
		},
	}
}


type Sujeto struct {
	Nombre			string				`bson:"nombre,omitempty"`
	Edad			int					`bson:"edad,omitempty"`
	Personeria		EnumPersoneria		`bson:"personeria,omitempty"`
	Nacionalidad	EnumNacionalidad	`bson:"nacionalidad,omitempty"`

	// Aspectos
	Operaciones		[]*Operacion		`bson:"actividades,omitempty"`
}
// constructor [init | gen]
func SujetoNew() *Sujeto {
	return &Sujeto {
		Nombre:			"Lucas",
		Edad:			23,
		Personeria:		0,
		Nacionalidad:	0,

		Operaciones:	[]*Operacion {
			OperacionNew(1),
			OperacionNew(2),
			OperacionNew(3),
		},
	}
}
// methods
func (s *Sujeto) is_underAge() bool {
	if s.Edad < 18 {
		return true
	}
	return false
}