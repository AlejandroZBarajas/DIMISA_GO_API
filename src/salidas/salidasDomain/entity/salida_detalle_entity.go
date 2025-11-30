package salidaEntity

type SalidaDetalleEntity struct {
	Id_salidaDetalle int32 `json:"id_salida_detalle"`
	Id_salida        int32 `json:"id_salida"`
	Id_medicamento   int32 `json:"id_medicamento"`
	Cantidad         int32 `json:"cantidad"`
}
