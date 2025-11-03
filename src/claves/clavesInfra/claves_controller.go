package clavesInfra

import (
	"DIMISA/src/claves/clavesApp"
	claveEntity "DIMISA/src/claves/clavesDomain/entity"
	"encoding/json"
	"net/http"
	"strings"
)

type ClavesController struct {
	SearchUC *clavesApp.SearchClave
}

func NewClaveController(search *clavesApp.SearchClave) *ClavesController {
	return &ClavesController{
		SearchUC: search,
	}
}

// Response structure
type SearchResponse struct {
	Success bool                       `json:"success"`
	Data    []*claveEntity.ClaveEntity `json:"data,omitempty"`
	Message string                     `json:"message,omitempty"`
	Count   int                        `json:"count"`
}

// SearchForClave maneja búsquedas con GET /medicamentos/search?q=query
func (c *ClavesController) SearchForClave(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Solo permitir GET
	if r.Method != http.MethodGet {
		sendError(w, "Método no permitido. Usa GET", http.StatusMethodNotAllowed)
		return
	}

	// Obtener query parameter
	query := r.URL.Query().Get("q")
	query = strings.TrimSpace(query)

	// Validar que no esté vacío
	if query == "" {
		sendError(w, "El parámetro 'q' es requerido", http.StatusBadRequest)
		return
	}

	// Validar longitud mínima (evitar búsquedas muy cortas)
	if len(query) < 2 {
		sendError(w, "La búsqueda debe tener al menos 2 caracteres", http.StatusBadRequest)
		return
	}

	// Ejecutar búsqueda
	results, err := c.SearchUC.Execute(query)
	if err != nil {
		sendError(w, "Error al buscar medicamentos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con resultados
	response := SearchResponse{
		Success: true,
		Data:    results,
		Count:   len(results),
	}

	// Si no hay resultados
	if len(results) == 0 {
		response.Message = "No se encontraron medicamentos"
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Helper para enviar errores
func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(SearchResponse{
		Success: false,
		Message: message,
		Count:   0,
	})
}
