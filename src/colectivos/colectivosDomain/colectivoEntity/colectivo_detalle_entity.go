package colectivoEntity

type ColectivoDetalleEntity struct {
	Id_detalle     int32  `json:"id_detalle"`
	Id_colectivo   int32  `json:"id_colectivo"`
	Id_medicamento int32  `json:"id_medicamento"`
	Clave          string `json:"clave"`
	Descripcion    string `json:"descripcion"`
	Cantidad       int32  `json:"cantidad"`
}
