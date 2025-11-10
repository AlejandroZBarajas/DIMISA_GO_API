package colectivoEntity

type ColectivoEntity struct {
	Id_colectivo int32                    `json:"id_colectivo"`
	Folio        string                   `json:"folio"`
	Fecha        string                   `json:"fecha"`
	Id_user      int32                    `json:"id_user"`
	Id_area      int32                    `json:"id_area"`
	Id_cendis    int32                    `json:"id_cendis"`
	Capturado    bool                     `json:"capturado"`
	Claves       []ColectivoDetalleEntity `json:"claves"`
}
