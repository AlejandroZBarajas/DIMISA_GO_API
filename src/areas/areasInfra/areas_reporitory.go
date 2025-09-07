package areasInfra

import (
	"DIMISA/src/areas/areasDomain"
	"DIMISA/src/areas/areasDomain/areaEntity"
	"database/sql"
	"fmt"
)

type AreasRepository struct {
	DB *sql.DB
}

func (r *AreasRepository) CreateArea(area *areaEntity.AreaEntity) error {
	query := "INSERT INTO areas (nombre_area, cama_1, cama_n) VALUES (?, ?, ?)"
	res, err := r.DB.Exec(query, area.Nombre_area, area.Cama_1, area.Cama_n)
	if err != nil {
		return fmt.Errorf("error al crear área: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener id_area generado: %w", err)
	}

	area.Id_area = int32(id)
	return nil
}

func (r *AreasRepository) UpdateArea(area *areaEntity.AreaEntity) error {
	query := "UPDATE areas SET nombre_area = ?, cama_1 = ?, cama_n = ? WHERE id_area = ?"
	_, err := r.DB.Exec(query, area.Nombre_area, area.Cama_1, area.Cama_n, area.Id_area)
	if err != nil {
		return fmt.Errorf("error al actualizar área: %w", err)
	}
	return nil
}

func (r *AreasRepository) GetAllAreas() ([]*areaEntity.AreaEntity, error) {
	query := "SELECT id_area, nombre_area, cama_1, cama_n FROM areas"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todas las áreas: %w", err)
	}
	defer rows.Close()

	var areas []*areaEntity.AreaEntity
	for rows.Next() {
		area := &areaEntity.AreaEntity{}
		if err := rows.Scan(&area.Id_area, &area.Nombre_area, &area.Cama_1, &area.Cama_n); err != nil {
			return nil, fmt.Errorf("error al escanear área: %w", err)
		}
		areas = append(areas, area)
	}

	return areas, nil
}

func (r *AreasRepository) GetAreaByID(id int32) (*areaEntity.AreaEntity, error) {
	query := "SELECT id_area, nombre_area, cama_1, cama_n FROM areas WHERE id_area = ?"
	row := r.DB.QueryRow(query, id)

	area := &areaEntity.AreaEntity{}
	if err := row.Scan(&area.Id_area, &area.Nombre_area, &area.Cama_1, &area.Cama_n); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al obtener área por ID: %w", err)
	}
	return area, nil
}

func (r *AreasRepository) DeleteArea(id int32) error {
	query := "DELETE FROM areas WHERE id_area = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar área: %w", err)
	}
	return nil
}

var _ areasDomain.AreasInterface = &AreasRepository{}
