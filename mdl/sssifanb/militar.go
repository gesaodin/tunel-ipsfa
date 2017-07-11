package sssifanb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gesaodin/tunel-ipsfa/sys"
	"github.com/gesaodin/tunel-ipsfa/util"
	"gopkg.in/mgo.v2/bson"
)

const (
	MILITAR int = 0
)

type Militar struct {
	//Persona                DatoBasico  `json:"Persona,omitempty" bson:"persona"`
	// Direccion              []Direccion `json:"Direccion,omitempty" bson:"direccion"`
	// Telefono               []Telefono  `json:"Telefono,omitempty" bson:"telefono"`a
	// Correo                 Correo      `json:"Correo,omitempty" bson:"correo"`
	ID                     int        `json:"id,omitempty" bson:"id"`
	TipoDato               int        `json:"tipodato,omitempty" bson:"tipodato"`
	Persona                Persona    `json:"Persona,omitempty" bson:"persona"`
	Categoria              string     `json:"Categoria,omitempty" bson:"categoria"` // efectivo,asimilado,invalidez, reserva activa, tropa
	Situacion              string     `json:"Situacion,omitempty" bson:"situacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	Clase                  int        `json:"clase,omitempty" bson:"clase"`         //alumno, cadete, oficial, oficial tecnico, oficial tropa, sub.oficial
	FechaIngresoComponente time.Time  `json:"fingreso,omitempty" bson:"fingreso"`
	FechaAscenso           time.Time  `json:"fascenso,omitempty" bson:"fascenso"`
	AnoReconocido          string     `json:"areconocido,omitempty" bson:"areconocido"`
	MesReconocido          string     `json:"mreconocido,omitempty" bson:"mreconocido"`
	DiaReconocido          string     `json:"dreconocido,omitempty" bson:"dreconocido"`
	NumeroResuelto         string     `json:"nresuelto,omitempty" bson:"nresuelto"`
	Posicion               int        `json:"posicion,omitempty" bson:"posicion"`
	DescripcionHistorica   string     `json:"dhistorica,omitempty" bson:"dhistorica"` //codigo
	Componente             Componente `json:"Componente,omitempty" bson:"componente"`
	Grado                  Grado      `json:"Grado,omitempty" bson:"grado"` //grado
	TIM                    Carnet     `json:"Tim,omitempty" bson:"tim"`     //Tarjeta de Identificacion Militar
	Familiar               []Familiar `json:"Familiar" bson:"familiar"`
	AppSaman               bool
	AppPace                bool
	AppNomina              bool
}

type HistorialMilitar struct {
	Categoria              string     `json:"Categoria,omitempty" bson:"categoria"` // efectivo,asimilado,invalidez, reserva activa, tropa
	Situacion              string     `json:"Situacion,omitempty" bson:"situacion"` //activo,fallecido con pension, fsp, retirado con pension, rsp
	Clase                  int        `json:"clase,omitempty" bson:"clase"`         //alumno, cadete, oficial, oficial tecnico, oficial tropa, sub.oficial
	FechaIngresoComponente time.Time  `json:"fingreso,omitempty" bson:"fingreso"`
	FechaAscenso           time.Time  `json:"fascenso,omitempty" bson:"fascenso"`
	AnoReconocido          string     `json:"areconocido,omitempty" bson:"areconocido"`
	MesReconocido          string     `json:"mreconocido,omitempty" bson:"mreconocido"`
	DiaReconocido          string     `json:"dreconocido,omitempty" bson:"dreconocido"`
	NumeroResuelto         string     `json:"nresuelto,omitempty" bson:"nresuelto"`
	Posicion               int        `json:"posicion,omitempty" bson:"posicion"`
	DescripcionHistorica   string     `json:"dhistorica,omitempty" bson:"dhistorica"` //codigo
	Componente             Componente `json:"Componente,omitempty" bson:"componente"`
	Grado                  Grado      `json:"Grado,omitempty" bson:"grado"` //grado
	TIM                    Carnet     `json:"Tim,omitempty" bson:"tim"`     //Tarjeta de Identificacion Militar
}

type Componente struct {
	ID          int    `json:"id" bson:"id"`
	Nombre      string `json:"nombre" bson:"nombre"`
	Descripcion string `json:"descripcion" bson:"descripcion"`
	Abreviatura string `json:"abreviatura" bson:"abreviatura"`
}

type Grado struct {
	ID          int    `json:"id" bson:"id"`
	Nombre      string `json:"nombre" bson:"nombre"`
	Descripcion string `json:"descripcion" bson:"descripcion"`
	Abreviatura string `json:"abreviatura" bson:"abreviatura"`
}

//
func (m *Militar) Listar() {
	//gesaodin@gmail.com
}

//Mensaje del sistema
type Mensaje struct {
	Mensaje string `json:"msj"`
	Tipo    int    `json:"tipo"`
	Pgsql   string `json:"pgsql,omitempty"`
}

//ConsultarMGO una persona mediante el metodo de MongoDB
func (m *Militar) Consultar() (jSon []byte, err error) {
	var militar Militar
	var msj Mensaje
	c := sys.MGOSession.DB("ipsfa_test").C("militar")
	err = c.Find(bson.M{"persona.datobasico.cedula": m.Persona.DatoBasico.Cedula}).One(&militar)
	if militar.Persona.DatoBasico.Cedula == "" {
		msj.Tipo = 0
		jSon, err = json.Marshal(msj)
	} else {
		jSon, err = json.Marshal(militar)
	}
	return
}

//Consultar Militar
func (m *Militar) ConsultarSAMAN() (jSon []byte, err error) {
	var msj Mensaje
	var lst []Militar
	var estatus bool
	s := `SELECT codnip,tipnip, nropersona,nombreprimero, nombresegundo,apellidoprimero,apellidosegundo,sexocod
	FROM personas
	WHERE codnip='` + m.Persona.DatoBasico.Cedula + `' AND tipnip != 'P'`
	sq, err := sys.PostgreSQLSAMAN.Query(s)
	if err != nil {
		msj.Mensaje = "Error: Consulta ya existe."
		msj.Tipo = 2
		msj.Pgsql = err.Error()
		jSon, err = json.Marshal(msj)
		fmt.Println(err.Error())
		return
	}
	estatus = true
	for sq.Next() {
		var m Militar
		var cedula, tipnip string
		var nombp, nombs, apellp, apells, sexo sql.NullString
		var numero int

		sq.Scan(&cedula, &tipnip, &numero, &nombp, &nombs, &apellp, &apells, &sexo)
		m.Persona.DatoBasico.Cedula = cedula
		m.Persona.DatoBasico.NumeroPersona = numero
		m.Persona.DatoBasico.NombrePrimero = util.ValidarNullString(nombp)
		m.Persona.DatoBasico.NombreSegundo = util.ValidarNullString(nombs)
		m.Persona.DatoBasico.ApellidoPrimero = util.ValidarNullString(apellp)
		m.Persona.DatoBasico.ApellidoSegundo = util.ValidarNullString(apells)
		m.Persona.DatoBasico.Nacionalidad = tipnip
		m.Persona.DatoBasico.Sexo = util.ValidarNullString(sexo)
		if m.Persona.DatoBasico.NombrePrimero != "null" {
			estatus = false
		} else {
			estatus = true
		}

		lst = append(lst, m)

	}
	if estatus == true {
		msj.Mensaje = "Afiliado no existe."
		msj.Tipo = 0
		jSon, err = json.Marshal(msj)
	} else {
		jSon, err = json.Marshal(lst)
	}

	return

}

//Actualizar Vida Militar
func (m *Militar) Actualizar() (jSon []byte, err error) {
	var msj Mensaje
	m.TipoDato = 0

	s := `UPDATE personas SET nombreprimero='` +
		m.Persona.DatoBasico.NombrePrimero +
		`', nombresegundo='` +
		m.Persona.DatoBasico.NombreSegundo +
		`' WHERE codnip='` + m.Persona.DatoBasico.Cedula + `'`
	_, err = sys.PostgreSQLSAMAN.Exec(s)
	if err != nil {
		msj.Mensaje = "Error: Consulta ya existe."
		msj.Tipo = 2
		msj.Pgsql = err.Error()
		jSon, err = json.Marshal(msj)
		return
	}
	msj.Mensaje = "Su data ha sido actualizada."
	msj.Tipo = 2
	jSon, err = json.Marshal(msj)
	m.SalvarMGO("")
	return
}

//ActualizarMGO Actualizar
func (m *Militar) ActualizarMGO(oid string, familiar map[string]interface{}) (err error) {
	c := sys.MGOSession.DB("ipsfa_test").C("militar")
	err = c.Update(bson.M{"persona.datobasico.cedula": oid}, bson.M{"$set": familiar})
	if err != nil {
		fmt.Println("Cedula: " + oid + " -> " + err.Error())
		return
	}
	return
}

//SalvarMGO Guardar
func (m *Militar) SalvarMGO(colecion string) (err error) {
	if colecion != "" {
		c := sys.MGOSession.DB("ipsfa_test").C(colecion)
		err = c.Insert(m)
	} else {
		c := sys.MGOSession.DB("ipsfa_test").C("persona")
		err = c.Insert(m)
	}

	//fmt.Println(p)

	return
}

//ConsultarMGO una persona mediante el metodo de MongoDB
func (m *Militar) ConsultarMGO(cedula string) (err error) {
	c := sys.MGOSession.DB("ipsfa_test").C("persona")
	err = c.Find(bson.M{"cedula": cedula}).One(&m)
	return
}

//ListarMGO Listado General
func (m *Militar) ListarMGO(cedula string) (lst []Militar, err error) {
	c := sys.MGOSession.DB("ipsfa_test").C("persona")
	err = c.Find(bson.M{}).All(&lst)
	return
}
