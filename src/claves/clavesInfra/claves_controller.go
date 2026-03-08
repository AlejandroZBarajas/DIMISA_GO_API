package clavesInfra

import (
	"DIMISA/src/claves/clavesApp"
	claveEntity "DIMISA/src/claves/clavesDomain/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ClavesController struct {
	SearchUC          *clavesApp.SearchClave
	SearchInventoryUC *clavesApp.SearchInInventory
}

func NewClaveController(
	search *clavesApp.SearchClave,
	searchInventory *clavesApp.SearchInInventory) *ClavesController {
	return &ClavesController{
		SearchUC:          search,
		SearchInventoryUC: searchInventory,
	}
}

type SearchResponse struct {
	Success bool                       `json:"success"`
	Data    []*claveEntity.ClaveEntity `json:"data,omitempty"`
	Message string                     `json:"message,omitempty"`
	Count   int                        `json:"count"`
}

// SearchForClave maneja búsquedas con GET /medicamentos/search?q=query
func (c *ClavesController) SearchForClave(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("entra al backend")
	if r.Method != http.MethodGet {
		sendError(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("q")
	query = strings.TrimSpace(query)

	if query == "" {
		sendError(w, "El parámetro 'q' es requerido", http.StatusBadRequest)
		return
	}

	if len(query) < 2 {
		sendError(w, "La búsqueda debe tener al menos 2 caracteres", http.StatusBadRequest)
		return
	}

	results, err := c.SearchUC.Execute(query)
	if err != nil {
		sendError(w, "Error al buscar medicamentos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := SearchResponse{
		Success: true,
		Data:    results,
		Count:   len(results),
	}

	if len(results) == 0 {
		response.Message = "No se encontraron medicamentos"
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(SearchResponse{
		Success: false,
		Message: message,
		Count:   0,
	})
}

func (c *ClavesController) SearchInInventory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		sendError(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		sendError(w, "El parámetro 'q' es requerido", http.StatusBadRequest)
		return
	}
	if len(query) < 2 {
		sendError(w, "La búsqueda debe tener al menos 2 caracteres", http.StatusBadRequest)
		return
	}

	cendisStr := r.URL.Query().Get("cendis")
	if cendisStr == "" {
		sendError(w, "El parámetro 'cendis' es requerido", http.StatusBadRequest)
		return
	}

	cendisID, err := strconv.Atoi(cendisStr)
	if err != nil || cendisID <= 0 {
		sendError(w, "El parámetro 'cendis' debe ser un número válido", http.StatusBadRequest)
		return
	}

	results, err := c.SearchInventoryUC.Execute(query, int32(cendisID))
	if err != nil {
		sendError(w, "Error al buscar en inventario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := SearchResponse{
		Success: true,
		Data:    results,
		Count:   len(results),
	}
	if len(results) == 0 {
		response.Message = "No se encontraron medicamentos en el inventario"
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
