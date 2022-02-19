package models

import (
	"fmt"
	"log"
	"encoding/json"

	"github.com/LNMMusic/fce/handlers"
)

// STRUCT
type IVAInstance struct {
	Cases		handlers.Map
}
// Methods
func (i *IVAInstance) Load() {
	bytes, err := handlers.FileToBytes("taxes.json")
	if err != nil {
		log.Fatalf("Failed to Parse File in Bytes -> %v", err)
	}

	var cases = handlers.Map{}
	if err := json.Unmarshal([]byte(bytes), &cases); err != nil {
		log.Fatalf("Failed to Parse Bytes in Struct -> %v", err)
	}

	i.Cases = cases.M("IVA")
}
func (i *IVAInstance) Resume(sujeto *Sujeto) {
	// Generates Codes from Subjet Data and Search in HashMap [based on 'personeria']
	var ivaSujeto = i.Cases.M(sujeto.Personeria.String())

	for _, operacion := range sujeto.Operaciones {
		c := fmt.Sprintf("op%d-ob%d-t%d",
			operacion.Tipo,
			operacion.Objeto,
			operacion.Titulo,
		)

		if alicuota, ok := ivaSujeto[c]; ok {
			operacion.Taxes.IVA = alicuota.(float64)
		}
	}
}


// INSTANCE
var IVA = new(IVAInstance)			//		== var IVA = &IVAInstance{}			or 		var IVA IVAInstance (witout pointer)
									//													its type would be a pointer unknowed (nill)
									//									https://yourbasic.org/golang/gotcha-nil-pointer-dereference/
