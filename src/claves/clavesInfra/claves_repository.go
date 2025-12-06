package clavesInfra

import (
	claveEntity "DIMISA/src/claves/clavesDomain/entity"
	"database/sql"
	"fmt"
)

type ClaveRepository struct {
	DB *sql.DB
}

func (r *ClaveRepository) SearchClave(s string) ([]*claveEntity.ClaveEntity, error) {
	// Búsqueda inteligente: busca en ambos campos con LIKE
	// Prioriza coincidencias exactas en clave, luego parciales en descripción
	query := `
		SELECT 
			id_medicamento, 
			clave_med, 
			descripcion
		FROM medicamentos
		WHERE 
			clave_med LIKE ? 
			OR descripcion LIKE ?
		ORDER BY 
			CASE 
				WHEN clave_med = ? THEN 1
				WHEN clave_med LIKE ? THEN 2
				ELSE 3
			END,
			clave_med
		LIMIT 50
	`

	// Preparar parámetros de búsqueda
	searchTerm := "%" + s + "%"
	exactMatch := s
	startsWithMatch := s + "%"

	rows, err := r.DB.Query(
		query,
		searchTerm,      // clave_med LIKE ?
		searchTerm,      // descripcion LIKE ?
		exactMatch,      // ORDER BY CASE exact match
		startsWithMatch, // ORDER BY CASE starts with
	)
	if err != nil {
		return nil, fmt.Errorf("error en query: %w", err)
	}
	defer rows.Close()

	claves := []*claveEntity.ClaveEntity{}

	for rows.Next() {
		var c claveEntity.ClaveEntity
		err := rows.Scan(
			&c.Id_medicamento,
			&c.Clave_med,
			&c.Descripcion,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear resultado: %w", err)
		}
		claves = append(claves, &c)
	}

	// Verificar errores durante la iteración
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar resultados: %w", err)
	}

	return claves, nil
}
