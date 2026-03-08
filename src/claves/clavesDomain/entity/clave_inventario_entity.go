package claveEntity

type ClaveInventarioEntity struct {
	Id_medicamento  int32  `json:"id_medicamento"`
	Clave_med       string `json:"clave_med"`
	Descripcion     string `json:"descripcion"`
	Cantidad_actual int32  `json:"cantidad_actual,omitempty"`
}
