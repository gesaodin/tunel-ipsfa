package sssifanb

type Direccion struct {
	Tipo         int    `json:"tipo,omitempty" bson:"tipo"` //domiciliaria, trabajo, emergencia
	Ciudad       string `json:"ciudad,omitempty" bson:"ciudad"`
	Estado       string `json:"estado,omitempty" bson:"estado"`
	Municipio    string `json:"municipio,omitempty" bson:"municipio"`
	Parroquia    string `json:"parroquia,omitempty" bson:"parroquia"`
	Calleavenida string `json:"calleavenida,omitempty" bson:"calleavenida"`
	Casa         string `json:"casa,omitempty" bson:"casa"`
	Apartamento  string `json:"apartamento,omitempty" bson:"apartamento"`
	Numero       int    `json:"numero,omitempty" bson:"numero"`
}
