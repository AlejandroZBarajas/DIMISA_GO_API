package cendisInfra

import (
	"DIMISA/src/cendis/cendisDomain"
	cendisEntity "DIMISA/src/cendis/cendisDomain/entity"
	"database/sql"
	"fmt"
)

type CendisRepository struct {
	DB *sql.DB
}

func (r *CendisRepository) CreateCendis(cendis *cendisEntity.CendisEntity, areaIDs []int32) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("error iniciando transacción: %w", err)
	}

	res, err := tx.Exec("INSERT INTO cendis (cendis_nombre) VALUES (?)", cendis.Cendis_nombre)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error al crear cendis: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error al obtener id_cendis generado: %w", err)
	}
	cendis.Id_cendis = int32(id)

	for _, areaID := range areaIDs {
		_, err := tx.Exec("INSERT INTO areas_cendis (id_area, id_cendis) VALUES (?, ?)", areaID, id)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error al asociar área %d: %w", areaID, err)
		}
	}
	return tx.Commit()
}

func (r *CendisRepository) UpdateCendis(cendis *cendisEntity.CendisEntity, areas []int32) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE cendis SET cendis_nombre = ? WHERE id_cendis = ?`,
		cendis.Cendis_nombre, cendis.Id_cendis)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`DELETE FROM areas_cendis WHERE id_cendis = ?`, cendis.Id_cendis)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, areaID := range areas {
		_, err = tx.Exec(`INSERT INTO areas_cendis (id_cendis, id_area) VALUES (?, ?)`,
			cendis.Id_cendis, areaID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *CendisRepository) GetAllCendis() ([]*cendisEntity.CendisEntity, error) {
	rows, err := r.DB.Query("SELECT id_cendis, cendis_nombre FROM cendis")
	if err != nil {
		return nil, fmt.Errorf("error al obtener todos los cendis: %w", err)
	}
	defer rows.Close()

	cendisList := []*cendisEntity.CendisEntity{}
	for rows.Next() {
		cendis := &cendisEntity.CendisEntity{}
		if err := rows.Scan(&cendis.Id_cendis, &cendis.Cendis_nombre); err != nil {
			return nil, fmt.Errorf("error al escanear cendis: %w", err)
		}
		cendisList = append(cendisList, cendis)
	}

	return cendisList, nil
}

func (r *CendisRepository) DeleteCendis(id int32) error {
	// eliminar relaciones primero
	_, err := r.DB.Exec("DELETE FROM areas_cendis WHERE id_cendis = ?", id)
	if err != nil {
		return fmt.Errorf("error al eliminar relaciones del cendis: %w", err)
	}

	// eliminar cendis
	_, err = r.DB.Exec("DELETE FROM cendis WHERE id_cendis = ?", id)
	if err != nil {
		return fmt.Errorf("error al eliminar cendis: %w", err)
	}

	return nil
}

var _ cendisDomain.CendisInterface = &CendisRepository{}
