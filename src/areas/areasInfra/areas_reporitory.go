package areasInfra

import (
	"DIMISA/src/areas/areasDomain/areaEntity"
	"database/sql"
	"fmt"
)

type AreasRepository struct {
	DB *sql.DB
}

func (r *AreasRepository) CreateArea(area *areaEntity.AreaEntity) error {
	query := "INSERT INTO areas (nombre_area, alias) VALUES (?, ?)"
	res, err := r.DB.Exec(query, area.Nombre_area, area.Alias)
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
	query := "UPDATE areas SET nombre_area = ?, alias = ? WHERE id_area = ?"
	_, err := r.DB.Exec(query, area.Nombre_area, area.Alias, area.Id_area)
	if err != nil {
		return fmt.Errorf("error API al actualizar área: %w", err)
	}
	return nil
}

func (r *AreasRepository) GetAllAreas() ([]*areaEntity.AreaEntity, error) {
	query := "SELECT id_area, nombre_area, alias FROM areas"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener todas las áreas: %w", err)
	}
	defer rows.Close()

	areas := []*areaEntity.AreaEntity{}
	for rows.Next() {
		area := &areaEntity.AreaEntity{}
		if err := rows.Scan(&area.Id_area, &area.Nombre_area, &area.Alias); err != nil {
			return nil, fmt.Errorf("error al escanear área: %w", err)
		}
		areas = append(areas, area)
	}

	return areas, nil
}

func (r *AreasRepository) GetAreaByID(id int32) (*areaEntity.AreaEntity, error) {
	query := "SELECT id_area, nombre_area, alias FROM areas WHERE id_area = ?"
	row := r.DB.QueryRow(query, id)

	area := &areaEntity.AreaEntity{}
	if err := row.Scan(&area.Id_area, &area.Nombre_area, &area.Alias); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error al obtener área por ID: %w", err)
	}
	return area, nil
}

func (r *AreasRepository) DeleteArea(id int32) error {
	query := "DELETE FROM camas WHERE id_area = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error en el back al eliminar las camas")
	}
	query = "DELETE FROM areas_cendis WHERE id_area = ?"
	_, err = r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el area de areas_cendis: %w", err)
	}

	query = "DELETE FROM areas WHERE id_area = ?"
	_, err = r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error en el back al eliminar área: %w", err)
	}
	return nil
}

func (r *AreasRepository) GetFreeAreas() ([]*areaEntity.AreaEntity, error) {
	query := `
		SELECT a.id_area, a.nombre_area, a.alias
		FROM areas a
		LEFT JOIN areas_cendis ac ON a.id_area = ac.id_area
		WHERE ac.id_area IS NULL
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener áreas libres: %w", err)
	}
	defer rows.Close()

	var areas []*areaEntity.AreaEntity
	for rows.Next() {
		area := &areaEntity.AreaEntity{}
		if err := rows.Scan(&area.Id_area, &area.Nombre_area, &area.Alias); err != nil {
			return nil, fmt.Errorf("error al escanear área libre: %w", err)
		}
		areas = append(areas, area)
	}

	return areas, nil
}

func (r *AreasRepository) GetAreasByCendis(id int32) ([]*areaEntity.AreaEntity, error) {
	query := `
	SELECT a.id_area, a.nombre_area, a.alias
	FROM areas a 
	INNER JOIN areas_cendis ac ON a.id_area = ac.id_area
	WHERE ac.id_cendis = ?
	`
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error al obtener áreas del cendis: %w", err)
	}
	defer rows.Close()

	var areas []*areaEntity.AreaEntity
	for rows.Next() {
		area := &areaEntity.AreaEntity{}
		if err := rows.Scan(&area.Id_area, &area.Nombre_area, &area.Alias); err != nil {
			return nil, fmt.Errorf("error al escanear área del cendis: %w", err)
		}
		areas = append(areas, area)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante la iteración de áreas: %w", err)
	}

	return areas, nil
}
