package salidasInfra

import (
	"DIMISA/src/salidas/salidasApp"
	salidaEntity "DIMISA/src/salidas/salidasDomain/entity"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type SalidasController struct {
	CreateUseCase               *salidasApp.CreateSalida
	UpdateUseCase               *salidasApp.UpdateSalida
	DeleteUseCase               *salidasApp.DeleteSalida
	GetSalidasByCendisUseCase   *salidasApp.GetSalidasByCendis
	GetSalidasPendientesUseCase *salidasApp.GetSalidasPendientes
}

func NewSalidasController(
	createUC *salidasApp.CreateSalida,
	updateUC *salidasApp.UpdateSalida,
	deleteUC *salidasApp.DeleteSalida,
	getSalidasByCendisUC *salidasApp.GetSalidasByCendis,
	getSalidasPendientesUC *salidasApp.GetSalidasPendientes,
) *SalidasController {
	return &SalidasController{
		CreateUseCase:               createUC,
		UpdateUseCase:               updateUC,
		GetSalidasByCendisUseCase:   getSalidasByCendisUC,
		DeleteUseCase:               deleteUC,
		GetSalidasPendientesUseCase: getSalidasPendientesUC,
	}
}

// POST /salidas
func (ctrl *SalidasController) CreateSalidaHnadler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var salida salidaEntity.SalidaEntity

	if err := json.NewDecoder(r.Body).Decode(&salida); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if err := ctrl.CreateUseCase.Execute(&salida); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Salida creada exitosamente"})
}

// PUT /salidas/{id}
func (ctrl *SalidasController) UpdateSalida(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Extraer ID de la URL
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 2 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID no proporcionado"})
		return
	}

	id, err := strconv.ParseInt(pathParts[len(pathParts)-1], 10, 32)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID inválido"})
		return
	}

	var salida salidaEntity.SalidaEntity

	if err := json.NewDecoder(r.Body).Decode(&salida); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	salida.Id_salida = int32(id)

	if err := ctrl.UpdateUseCase.Execute(&salida); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Salida actualizada exitosamente"})
}

// GET /salidas/cendis/{id_cendis}
func (ctrl *SalidasController) GetSalidasByCendis(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Extraer ID de la URL
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 3 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID de cendis no proporcionado"})
		return
	}

	id_cendis, err := strconv.ParseInt(pathParts[len(pathParts)-1], 10, 32)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID de cendis inválido"})
		return
	}

	salidas, err := ctrl.GetSalidasByCendisUseCase.Execute(int32(id_cendis))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(salidas)
}

// DELETE /salidas/{id}
func (ctrl *SalidasController) DeleteSalida(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Extraer ID de la URL
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 2 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID no proporcionado"})
		return
	}

	id, err := strconv.ParseInt(pathParts[len(pathParts)-1], 10, 32)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID inválido"})
		return
	}

	if err := ctrl.DeleteUseCase.Execute(int32(id)); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Salida eliminada exitosamente"})
}
