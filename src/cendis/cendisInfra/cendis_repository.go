package cendisInfra

import (
	areaEntity "DIMISA/src/areas/areasDomain/areaEntity"
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

func (r *CendisRepository) DeleteCendis(id int32) error {
	println("primero se eliminan relaciones del cendis")
	_, err := r.DB.Exec("DELETE FROM areas_cendis WHERE id_cendis = ?", id)
	if err != nil {
		return fmt.Errorf("error al eliminar relaciones del cendis: %w", err)
	}

	println("y luego se elimina el cendis")
	_, err = r.DB.Exec("DELETE FROM cendis WHERE id_cendis = ?", id)
	if err != nil {
		return fmt.Errorf("error al eliminar cendis: %w", err)
	}

	return nil
}

func (r *CendisRepository) GetAllCendis() ([]*cendisEntity.CendisWithAreas, error) {
	query := `
		SELECT 
			c.id_cendis, 
			c.cendis_nombre,
			a.id_area,
			a.nombre_area,
			a.alias
		FROM cendis c
		LEFT JOIN areas_cendis ac ON c.id_cendis = ac.id_cendis
		LEFT JOIN areas a ON ac.id_area = a.id_area
		ORDER BY c.id_cendis;
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cendis con áreas: %w", err)
	}
	defer rows.Close()

	cendisMap := make(map[int32]*cendisEntity.CendisWithAreas)

	for rows.Next() {
		var (
			idCendis     int32
			nombreCendis string
			idArea       sql.NullInt32
			nombreArea   sql.NullString
			alias        sql.NullString
		)

		if err := rows.Scan(&idCendis, &nombreCendis, &idArea, &nombreArea, &alias); err != nil {
			return nil, fmt.Errorf("error al escanear fila: %w", err)
		}

		// Si el cendis no está en el mapa, lo agregamos
		if _, exists := cendisMap[idCendis]; !exists {
			cendisMap[idCendis] = &cendisEntity.CendisWithAreas{
				Id_cendis:     idCendis,
				Cendis_nombre: nombreCendis,
				Areas:         []areaEntity.AreaEntity{},
			}
		}

		// Si el área no es nula, la agregamos
		if idArea.Valid {
			cendisMap[idCendis].Areas = append(
				cendisMap[idCendis].Areas,
				areaEntity.AreaEntity{
					Id_area:     idArea.Int32,
					Nombre_area: nombreArea.String,
					Alias:       alias.String,
				},
			)
		}
	}

	// Convertir el mapa a slice
	cendisList := make([]*cendisEntity.CendisWithAreas, 0, len(cendisMap))
	for _, c := range cendisMap {
		cendisList = append(cendisList, c)
	}

	return cendisList, nil
}

var _ cendisDomain.CendisInterface = &CendisRepository{}
