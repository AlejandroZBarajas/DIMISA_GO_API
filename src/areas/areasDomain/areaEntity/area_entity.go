package areaEntity

type AreaEntity struct {
	Id_area     int32  `json:"id_area"`
	Nombre_area string `json:"nombre_area"`
	Cama_1      int32  `json:"cama_1"`
	Cama_n      int32  `json:"cama_n"`
}
