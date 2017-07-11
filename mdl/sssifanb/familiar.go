package sssifanb

import "github.com/gesaodin/tunel-ipsfa/util"

type Familiar struct {
	ID         int     `json:"id" bson:"id"`
	Persona    Persona `json:"Persona" bson:"persona"`
	Parentesco string  `json:"parentesco" bson:"parentesco"` //0:Mama, 1:papa, 2: Esposa  3: hijo
	EsMilitar  bool    `json:"esmilitar" bson:"esmilitar"`
	Condicion  int     `json:"condicion" bson:"condicion"` //Sano o Condicion especial
	Estudia    int     `json:"estudia" bson:"estudia"`
	Benficio   bool    `json:"beneficio" bson:"beneficio"` //
	Documento  int     `json:"documento" bson:"documento"`
	Adoptado   bool    `json:"adoptado" bson:"adoptado"`
	//DocumentoPadre string
}

//AplicarReglasBeneficio OJO SEGUROS HORIZONTES
func (f *Familiar) AplicarReglasBeneficio() {
	if f.Parentesco == "HJ" {
		f.Benficio = false
		edad := util.CalcularEdad(f.Persona.DatoBasico.FechaNacimiento)
		if f.Condicion == 1 {
			f.Benficio = true
		} else {
			if edad < 18 {
				f.Benficio = true
			} else if f.Estudia == 1 && edad < 27 {
				f.Benficio = true
			}
		}
	} else { // ESPOSA Y PADRES
		f.Benficio = true
	}

}

func (f *Familiar) AplicarReglasParentesco() {

}
