package tiposInfra

import (
	tipoEntity "DIMISA/src/tipos_colectivo_salida/tiposDomain/entity"
	"database/sql"
)

type TiposRepository struct {
	DB *sql.DB
}

func NewTiposRepository(db *sql.DB) *TiposRepository {
	return &TiposRepository{DB: db}
}

func (r *TiposRepository) GetTipos() (*[]tipoEntity.TipoEntity, error) {
	query := `
		SELECT id_tipo, nombre 
		FROM tipos
		ORDER BY id_tipo
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tipos []tipoEntity.TipoEntity
	for rows.Next() {
		var t tipoEntity.TipoEntity
		if err := rows.Scan(&t.Id_tipo, &t.Nombre); err != nil {
			return nil, err
		}
		tipos = append(tipos, t)
	}

	return &tipos, nil
}
