package colectivosInfra

import (
	"DIMISA/src/colectivos/colectivosDomain/colectivoEntity"
	"database/sql"
	"fmt"
)

type ColectivoRepository struct {
	DB *sql.DB
}

func NewColectivoRepository(db *sql.DB) *ColectivoRepository {
	return &ColectivoRepository{DB: db}
}
func (r *ColectivoRepository) CreateColectivo(colectivo *colectivoEntity.ColectivoEntity) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	var idArea interface{}
	if colectivo.Id_area == 0 {
		idArea = nil
	} else {
		idArea = colectivo.Id_area
	}

	queryColectivo := `
		INSERT INTO colectivos (folio, fecha, id_user, id_area, id_cendis, capturado)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := tx.Exec(queryColectivo,
		nil, // folio se actualiza después
		colectivo.Fecha,
		colectivo.Id_user,
		idArea, // aquí ya puede ser NULL si aplica
		colectivo.Id_cendis,
		colectivo.Capturado,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	colectivo.Id_colectivo = int32(id)

	folio := fmt.Sprintf("F-%d", id)
	colectivo.Folio = folio
	_, err = tx.Exec(`UPDATE colectivos SET folio = ? WHERE id_colectivo = ?`, folio, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryDetalle := `
		INSERT INTO colectivo_detalle (id_colectivo, id_medicamento, cantidad)
		VALUES (?, ?, ?)
		`
	fmt.Printf("🧩 Claves recibidas: %+v\n", colectivo.Claves)
	for _, d := range colectivo.Claves {

		_, err := tx.Exec(queryDetalle, id, d.Id_medicamento, d.Cantidad)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *ColectivoRepository) GetColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoEntity, error) {
	query := `
		SELECT id_colectivo, folio, fecha, id_user, id_area, id_cendis, capturado
		FROM colectivos
		WHERE id_cendis = ?
	`
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var colectivos []*colectivoEntity.ColectivoEntity
	for rows.Next() {
		var c colectivoEntity.ColectivoEntity
		if err := rows.Scan(
			&c.Id_colectivo,
			&c.Folio,
			&c.Fecha,
			&c.Id_user,
			&c.Id_area,
			&c.Id_cendis,
			&c.Capturado,
		); err != nil {
			return nil, err
		}
		colectivos = append(colectivos, &c)
	}
	return colectivos, nil
}

func (r *ColectivoRepository) GetPendingColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoEntity, error) {
	query := `
		SELECT id_colectivo, folio, fecha, id_user, id_area, id_cendis, capturado
		FROM colectivos
		WHERE id_cendis = ? AND capturado = 0
	`
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pendientes []*colectivoEntity.ColectivoEntity
	for rows.Next() {
		var c colectivoEntity.ColectivoEntity
		if err := rows.Scan(
			&c.Id_colectivo,
			&c.Folio,
			&c.Fecha,
			&c.Id_user,
			&c.Id_area,
			&c.Id_cendis,
			&c.Capturado,
		); err != nil {
			return nil, err
		}
		pendientes = append(pendientes, &c)
	}
	return pendientes, nil
}
