package models

type Enums interface {
	String() string
}
// Enums [const block -> each const represent an 'enum...' type with some specific value
//		 the value assigned can match with the actual value assigned to the actual 'enum...' type
//		 and return something if its match, in this case an string]

// SUJETO
type EnumPersoneria int
const (
	Humana		EnumPersoneria = iota
	Juridica
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

type EnumNacionalidad int
const (
	Nativo		EnumNacionalidad = iota
	Extranjero
)
func (e EnumNacionalidad) String() string {
	switch e {
	case Nativo:
		return "Nativo"
	case Extranjero:
		return "Extranjero"
	default:
		return ""
	}
}



// OPERACION
type EnumOperacion int
const (
	Venta		EnumOperacion = iota
	Compra
	Permuta
	Importacion
	Exportacion
)
func (e EnumOperacion) String() string {
	switch e {
	case Venta:
		return "Venta"
	case Compra:
		return "Compra"
	case Permuta:
		return "Permuta"
	case Importacion:
		return "Importacion"
	case Exportacion:
		return "Exportacion"
	default:
		return ""
	}
}

type EnumObjeto int
const (
	BienMueble	EnumObjeto = iota
	BienInmueble
	Servicio
	Other
)
func (e EnumObjeto) String() string {
	switch e {
	case BienMueble:
		return "BienMueble"
	case BienInmueble:
		return "BienInmueble"
	case Servicio:
		return "Servicio"
	case Other:
		return "Other"
	default:
		return ""
	}
}

type EnumTitulo int
const (
	Oneroso		EnumTitulo = iota
	Gratuito
)
func (e EnumTitulo) String() string {
	switch e {
	case Oneroso:
		return "Oneroso"
	case Gratuito:
		return "Gratuito"
	default:
		return ""
	}
}

