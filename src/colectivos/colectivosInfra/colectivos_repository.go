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
		INSERT INTO colectivos (tipo_id, fecha, id_user, id_area, id_cendis)
		VALUES (?, ?, ?, ?, ?)
	`
	result, err := tx.Exec(queryColectivo,
		colectivo.Tipo_id,
		colectivo.Fecha,
		colectivo.Id_user,
		colectivo.Id_area,
		colectivo.Id_cendis,
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

func (r *ColectivoRepository) GetColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoDTO, error) {
	query := `
		SELECT 
			c.id_colectivo,
			CONCAT('F-', c.id_colectivo) AS folio,
			c.tipo_id,
			t.nombre AS tipo,
			c.fecha,
			c.id_user,
			CONCAT(u.nombres, ' ', u.apellido1, ' ', u.apellido2) AS nombre_usuario,
			c.id_area,
			c.id_cendis,
			ce.cendis_nombre AS cendis
		FROM colectivos c
		INNER JOIN tipos t ON c.tipo_id = t.id_tipo
		INNER JOIN usuarios u ON c.id_user = u.id_usuario
		INNER JOIN cendis ce ON c.id_cendis = ce.id_cendis
		WHERE c.id_cendis = ?
	`
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var colectivos []*colectivoEntity.ColectivoDTO
	for rows.Next() {
		var c colectivoEntity.ColectivoDTO
		if err := rows.Scan(
			&c.Id_colectivo,   // 1
			&c.Folio,          // 2
			&c.Tipo_id,        // 3
			&c.Tipo,           // 4
			&c.Fecha,          // 5
			&c.Id_user,        // 6
			&c.Nombre_usuario, // 7
			&c.Id_area,        // 8
			&c.Id_cendis,      // 9
			&c.Cendis,         // 10
		); err != nil {
			return nil, err
		}
		colectivos = append(colectivos, &c)
	}
	return colectivos, nil
}

func (r *ColectivoRepository) getDetallesByColectivoID(id int32) ([]colectivoEntity.ColectivoDetalleDTO, error) {
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

	var detalles []colectivoEntity.ColectivoDetalleDTO

	for rows.Next() {
		var d colectivoEntity.ColectivoDetalleDTO
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

func (r *ColectivoRepository) GetPendingColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoDTO, error) {
	query := `
		SELECT 
			c.id_colectivo,
			CONCAT('F-', c.id_colectivo) AS folio,
			c.tipo_id,
			t.nombre AS tipo,
			c.fecha,
			c.id_user,
			CONCAT(u.nombres, ' ', u.apellido1, ' ', u.apellido2) AS nombre_usuario,
			c.id_area,
			c.id_cendis,
			ce.cendis_nombre AS cendis
		FROM colectivos c
		INNER JOIN tipos t ON c.tipo_id = t.id_tipo
		INNER JOIN usuarios u ON c.id_user = u.id_usuario
		INNER JOIN cendis ce ON c.id_cendis = ce.id_cendis
		WHERE c.id_cendis = ? AND c.capturado = 0
	`
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pendientes []*colectivoEntity.ColectivoDTO

	for rows.Next() {
		var c colectivoEntity.ColectivoDTO

		if err := rows.Scan(
			&c.Id_colectivo,   // 1
			&c.Folio,          // 2
			&c.Tipo_id,        // 3
			&c.Tipo,           // 4
			&c.Fecha,          // 5
			&c.Id_user,        // 6
			&c.Nombre_usuario, // 7
			&c.Id_area,        // 8
			&c.Id_cendis,      // 9
			&c.Cendis,         // 10
		); err != nil {
			return nil, err
		}
		detalles, err := r.getDetallesByColectivoID(c.Id_colectivo)
		if err != nil {
			return nil, err
		}
		c.Claves = detalles
		pendientes = append(pendientes, &c)
	}
	return pendientes, nil
}

func (r *ColectivoRepository) GetUpdatableColectivosByCendis(id int32) ([]*colectivoEntity.ColectivoDTO, error) {
	query := `
        SELECT 
            c.id_colectivo,
			CONCAT('F-', c.id_colectivo) AS folio,
			c.tipo_id,
			t.nombre AS tipo,
            c.fecha, 
            c.id_user,
            CONCAT(u.nombres, ' ', u.apellido1, ' ', u.apellido2) AS nombre_usuario,
            c.id_area, 
            c.id_cendis, 
			ce.cendis_nombre AS cendis
        FROM colectivos c
		INNER JOIN tipos t ON c.tipo_id = t.id_tipo
        INNER JOIN usuarios u ON c.id_user = u.id_usuario
		INNER JOIN cendis ce ON c.id_cendis = ce.id_cendis
        WHERE c.id_cendis = ? AND c.editable = 1
    `
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var editables []*colectivoEntity.ColectivoDTO
	for rows.Next() {
		var c colectivoEntity.ColectivoDTO
		if err := rows.Scan(
			&c.Id_colectivo,   // 1
			&c.Folio,          // 2
			&c.Tipo_id,        // 3
			&c.Tipo,           // 4
			&c.Fecha,          // 5
			&c.Id_user,        // 6
			&c.Nombre_usuario, // 7
			&c.Id_area,        // 8
			&c.Id_cendis,      // 9
			&c.Cendis,         // 10
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
