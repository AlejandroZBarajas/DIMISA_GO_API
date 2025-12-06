package salidaEntity

type SalidaEntity struct {
	Id_salida  int32                 `json:"id_salida"`
	Id_area    int32                 `json:"id_area"`
	Id_cendis  int32                 `json:"id_cendis"`
	Id_usuario int32                 `json:"id_usuario"`
	Fecha      string                `json:"fecha"`
	Created_at string                `json:"created_at"`
	Editable   int32                 `json:"editable"`
	Pendiente  int32                 `json:"pendiente"`
	Claves     []SalidaDetalleEntity `json:"claves"`
}
