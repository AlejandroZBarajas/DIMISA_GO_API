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

	queryColectivo := `
		INSERT INTO colectivos (folio, fecha, id_user, id_area, id_cendis, capturado)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := tx.Exec(queryColectivo,
		nil, // folio se actualiza después
		colectivo.Fecha,
		colectivo.Id_user,
		colectivo.Id_area,
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

func (r *ColectivoRepository) getDetallesByColectivoID(id int32) ([]colectivoEntity.ColectivoDetalleEntity, error) {
	query := `
		SELECT 
			cd.id_detalle,
			cd.id_colectivo,
			cd.id_medicamento,
			m.clave_med,
			m.descripcion,
			cd.cantidad
		FROM colectivo_detalle cd
		INNER JOIN medicamentos m ON m.id_medicamento = cd.id_medicamento
		WHERE cd.id_colectivo = ?;
	`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var detalles []colectivoEntity.ColectivoDetalleEntity

	for rows.Next() {
		var d colectivoEntity.ColectivoDetalleEntity
		if err := rows.Scan(
			&d.Id_detalle,
			&d.Id_colectivo,
			&d.Id_medicamento,
			&d.Clave,
			&d.Descripcion,
			&d.Cantidad,
		); err != nil {
			return nil, err
		}

		detalles = append(detalles, d)
	}

	return detalles, nil
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

		// 🔥 Aquí se agregan los detalles del colectivo
		detalles, err := r.getDetallesByColectivoID(c.Id_colectivo)
		if err != nil {
			return nil, err
		}
		c.Claves = detalles

		pendientes = append(pendientes, &c)
	}

	return pendientes, nil
}

func (r *ColectivoRepository) GetUpdatableColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoEntity, error) {
	query := `
        SELECT 
            c.id_colectivo, 
            c.folio, 
            c.fecha, 
            c.id_user,
            CONCAT(u.nombres, ' ', u.apellido1, ' ', u.apellido2) AS nombre_usuario,
            c.id_area, 
            c.id_cendis, 
            c.capturado
        FROM colectivos c
        INNER JOIN usuarios u ON c.id_user = u.id_usuario
        WHERE c.id_cendis = ? AND c.editable = 1
    `
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var editables []*colectivoEntity.ColectivoEntity
	for rows.Next() {
		var c colectivoEntity.ColectivoEntity
		if err := rows.Scan(
			&c.Id_colectivo,
			&c.Folio,
			&c.Fecha,
			&c.Id_user,
			&c.Nombre_usuario,
			&c.Id_area,
			&c.Id_cendis,
			&c.Capturado,
		); err != nil {
			return nil, err
		}

		detalles, err := r.getDetallesByColectivoID(c.Id_colectivo)
		if err != nil {
			return nil, err
		}
		c.Claves = detalles

		editables = append(editables, &c)
	}
	return editables, nil
}
