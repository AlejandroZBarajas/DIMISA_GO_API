package salidasInfra

import (
	salidaEntity "DIMISA/src/salidas/salidasDomain/entity"
	"database/sql"
	"fmt"
)

type SalidasRepository struct {
	DB *sql.DB
}

func NewSalidasRepository(db *sql.DB) *SalidasRepository {
	return &SalidasRepository{DB: db}
}

func (repo *SalidasRepository) CreateSalida(salida *salidaEntity.SalidaEntity) error {
	// Validar que el array de claves no esté vacío
	if len(salida.Claves) == 0 {
		return fmt.Errorf("la salida debe contener al menos un medicamento")
	}

	// Iniciar transacción
	tx, err := repo.DB.Begin()
	if err != nil {
		return fmt.Errorf("error al iniciar transacción: %w", err)
	}

	// Defer para manejar rollback en caso de error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insertar salida principal
	querySalida := `
		INSERT INTO salidas (id_area, id_cendis, id_usuario, fecha, editable, pendiente, created_at)
		VALUES (?, ?, ?, ?, ?, ?, NOW())
	`

	result, err := tx.Exec(
		querySalida,
		salida.Id_area,
		salida.Id_cendis,
		salida.Id_usuario,
		salida.Fecha,
		salida.Editable,
		salida.Pendiente,
	)

	if err != nil {
		return fmt.Errorf("error al crear salida: %w", err)
	}

	// Obtener el ID generado
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener id de salida: %w", err)
	}

	salida.Id_salida = int32(id)

	// Insertar detalles
	queryDetalle := `
		INSERT INTO salidas_detalle (id_salida, id_medicamento, cantidad)
		VALUES (?, ?, ?)
	`

	for _, detalle := range salida.Claves {
		_, err = tx.Exec(
			queryDetalle,
			salida.Id_salida,
			detalle.Id_medicamento,
			detalle.Cantidad,
		)

		if err != nil {
			return fmt.Errorf("error al insertar detalle de salida: %w", err)
		}
	}

	// Commit de la transacción
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error al confirmar transacción: %w", err)
	}

	return nil
}

func (repo *SalidasRepository) UpdateSalida(salida *salidaEntity.SalidaEntity) error {
	// Validar que el array de claves no esté vacío
	if len(salida.Claves) == 0 {
		return fmt.Errorf("la salida debe contener al menos un medicamento")
	}

	// Iniciar transacción
	tx, err := repo.DB.Begin()
	if err != nil {
		return fmt.Errorf("error al iniciar transacción: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Actualizar salida principal
	querySalida := `
		UPDATE salidas
		SET id_area = ?, id_cendis = ?, id_usuario = ?, fecha = ?, editable = ?, pendiente = ?
		WHERE id_salida = ?
	`

	result, err := tx.Exec(
		querySalida,
		salida.Id_area,
		salida.Id_cendis,
		salida.Id_usuario,
		salida.Fecha,
		salida.Editable,
		salida.Pendiente,
		salida.Id_salida,
	)

	if err != nil {
		return fmt.Errorf("error al actualizar salida: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar actualización: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("salida no encontrada")
	}

	// Eliminar todos los detalles anteriores
	queryDeleteDetalles := `DELETE FROM salidas_detalle WHERE id_salida = ?`
	_, err = tx.Exec(queryDeleteDetalles, salida.Id_salida)
	if err != nil {
		return fmt.Errorf("error al eliminar detalles anteriores: %w", err)
	}

	// Insertar nuevos detalles
	queryDetalle := `
		INSERT INTO salidas_detalle (id_salida, id_medicamento, cantidad)
		VALUES (?, ?, ?)
	`

	for _, detalle := range salida.Claves {
		_, err = tx.Exec(
			queryDetalle,
			salida.Id_salida,
			detalle.Id_medicamento,
			detalle.Cantidad,
		)

		if err != nil {
			return fmt.Errorf("error al insertar detalle de salida: %w", err)
		}
	}

	// Commit de la transacción
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error al confirmar transacción: %w", err)
	}

	return nil
}

func (repo *SalidasRepository) GetSalidasByCendis(id_cendis int32) (*[]salidaEntity.SalidaEntity, error) {
	query := `
		SELECT id_salida, id_area, id_cendis, id_usuario, fecha, created_at
		FROM salidas
		WHERE id_cendis = $1
		ORDER BY created_at DESC
	`

	rows, err := repo.DB.Query(query, id_cendis)
	if err != nil {
		return nil, fmt.Errorf("error al consultar salidas: %w", err)
	}
	defer rows.Close()

	var salidas []salidaEntity.SalidaEntity

	for rows.Next() {
		var salida salidaEntity.SalidaEntity
		err := rows.Scan(
			&salida.Id_salida,
			&salida.Id_area,
			&salida.Id_cendis,
			&salida.Id_usuario,
			&salida.Fecha,
			&salida.Created_at,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear salida: %w", err)
		}
		salidas = append(salidas, salida)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar salidas: %w", err)
	}

	return &salidas, nil
}

func (repo *SalidasRepository) DeleteSalida(id_salida int32) error {
	query := `DELETE FROM salidas WHERE id_salida = $1`

	result, err := repo.DB.Exec(query, id_salida)
	if err != nil {
		return fmt.Errorf("error al eliminar salida: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al verificar eliminación: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("salida no encontrada")
	}

	return nil
}

func (repo *SalidasRepository) GetSalidasPendientes(id_cendis int32) (*[]salidaEntity.SalidaEntity, error) {
	// Este método lo trabajamos después con más detalle
	return nil, fmt.Errorf("método no implementado aún")
}
