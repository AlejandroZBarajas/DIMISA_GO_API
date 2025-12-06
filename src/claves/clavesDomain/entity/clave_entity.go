package claveEntity

type ClaveEntity struct {
	Id_medicamento int32  `json:"id_medicamento"`
	Clave_med      string `json:"clave_med"`
	Descripcion    string `json:"descripcion"`
}
