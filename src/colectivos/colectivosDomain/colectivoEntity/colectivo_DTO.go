package colectivoEntity

import "database/sql"

type ColectivoDTO struct {
	Id_colectivo   int32                 `json:"id_colectivo"`
	Tipo_id        int32                 `json:"tipo_id"`
	Tipo           string                `json:"tipo"`
	Folio          string                `json:"folio"`
	Fecha          string                `json:"fecha"`
	Id_user        int32                 `json:"id_user"`
	Nombre_usuario string                `json:"nombre_usuario"`
	Id_area        sql.NullInt32         `json:"id_area"`
	Id_cendis      int32                 `json:"id_cendis"`
	Cendis         string                `json:"cendis"`
	Claves         []ColectivoDetalleDTO `json:"claves"`
}
