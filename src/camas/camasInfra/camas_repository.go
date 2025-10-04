package camasInfra

import (
	"DIMISA/src/camas/camasDomain/camaEntity"
	"database/sql"
)

type CamaRepository struct {
	DB *sql.DB
}

func (r *CamaRepository) CreateCama(cama *camaEntity.CamaEntity) error {
	query := `INSERT INTO camas (id_area, numero_cama)
	          VALUES (?, ?)`
	_, err := r.DB.Exec(query,
		cama.Id_area,
		cama.Numero_cama,
	)
	return err
}

func (r *CamaRepository) UpdateCama(cama *camaEntity.CamaEntity) error {
	query := `UPDATE camas SET 
	          id_area = ?, numero_cama = ?, nombres = ?, apellido1 = ?, apellido2 = ?, fecha_nac = ?, expediente = ?, riesgo_caida = ?, riesgo_ulcera = ?, habilitada = ?
	          WHERE id_cama = ?`
	_, err := r.DB.Exec(query,
		cama.Id_area,
		cama.Numero_cama,
		cama.Nombres,
		cama.Apellido1,
		cama.Apellido2,
		cama.Fecha_nac,
		cama.Expediente,
		cama.Riesgo_caida,
		cama.Riesgo_ulcera,
		cama.Habilitada,
		cama.Id_cama,
	)
	return err
}

func (r *CamaRepository) GetCamasByArea(areaid int32) ([]*camaEntity.CamaEntity, error) {
	query := `SELECT id_cama, id_area, numero_cama, nombres, apellido1, apellido2, fecha_nac, expediente, riesgo_caida, riesgo_ulcera, habilitada
	          FROM camas WHERE id_area = ?`
	rows, err := r.DB.Query(query, areaid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	camas := []*camaEntity.CamaEntity{}
	for rows.Next() {
		var c camaEntity.CamaEntity
		var nombres, apellido1, apellido2, fecha_nac, expediente, riesgo_caida, riesgo_ulcera sql.NullString
		var habilitadaInt int

		err := rows.Scan(
			&c.Id_cama,
			&c.Id_area,
			&c.Numero_cama,
			&nombres,
			&apellido1,
			&apellido2,
			&fecha_nac,
			&expediente,
			&riesgo_caida,
			&riesgo_ulcera,
			&habilitadaInt,
		)
		if err != nil {
			return nil, err
		}

		// Asignar valores, reemplazando NULL por cadena vacía
		c.Nombres = nombres.String
		c.Apellido1 = apellido1.String
		c.Apellido2 = apellido2.String
		c.Fecha_nac = fecha_nac.String
		c.Expediente = expediente.String
		c.Riesgo_caida = riesgo_caida.String
		c.Riesgo_ulcera = riesgo_ulcera.String
		c.Habilitada = habilitadaInt != 0

		camas = append(camas, &c)
	}

	return camas, nil
}

func (r *CamaRepository) EnableCama(id int32) error {
	query := `UPDATE camas SET habilitada = 1 WHERE id_cama = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *CamaRepository) DisableCama(id int32) error {
	query := `UPDATE camas SET habilitada = 0 WHERE id_cama = ?`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *CamaRepository) DeleteCama(id int32) error {
	query := `DELETE FROM camas WHERE id_cama = ?`
	_, err := r.DB.Exec(query, id)
	return err
}
