package camaEntity

type CamaEntity struct {
	Id_cama       int32  `json:"id_cama"`
	Id_area       int32  `json:"id_area"`
	Numero_cama   int32  `json:"numero_cama"`
	Nombres       string `json:"nombres"`
	Apellido1     string `json:"apellido1"`
	Apellido2     string `json:"apellido2"`
	Fecha_nac     string `json:"fecha_nac"`
	Expediente    string `json:"expediente"`
	Riesgo_caida  string `json:"riesgo_caida"`
	Riesgo_ulcera string `json:"riesgo_ulcera"`
	Habilitada    bool   `json:"habilitada"`
}
