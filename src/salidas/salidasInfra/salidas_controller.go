package salidasInfra

import (
	"DIMISA/src/salidas/salidasApp"
	salidaEntity "DIMISA/src/salidas/salidasDomain/entity"
	"encoding/json"
	"fmt"
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
	AddToSalidaUseCase          *salidasApp.AddToSalida
	CerrarSalidaUseCase         *salidasApp.CerrarSalida
}

func NewSalidasController(
	createUC *salidasApp.CreateSalida,
	updateUC *salidasApp.UpdateSalida,
	deleteUC *salidasApp.DeleteSalida,
	getSalidasByCendisUC *salidasApp.GetSalidasByCendis,
	getSalidasPendientesUC *salidasApp.GetSalidasPendientes,
	addToSalidaUC *salidasApp.AddToSalida,
	cerrarSalidaUC *salidasApp.CerrarSalida,
) *SalidasController {
	return &SalidasController{
		CreateUseCase:               createUC,
		UpdateUseCase:               updateUC,
		GetSalidasByCendisUseCase:   getSalidasByCendisUC,
		DeleteUseCase:               deleteUC,
		GetSalidasPendientesUseCase: getSalidasPendientesUC,
		AddToSalidaUseCase:          addToSalidaUC,
		CerrarSalidaUseCase:         cerrarSalidaUC,
	}
}

func (ctrl *SalidasController) CreateSalidaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entra al controller")
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

	// Ejecutar y capturar el ID retornado
	idSalida, err := ctrl.CreateUseCase.Execute(&salida)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Retornar el ID de la salida creada
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   "Salida creada exitosamente",
		"id_salida": idSalida,
	})
}

func (ctrl *SalidasController) UpdateSalidaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

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

	// CAMBIO: Pasar ambos parámetros (id del path y la entidad)
	if err := ctrl.UpdateUseCase.Execute(int32(id), &salida); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Salida actualizada exitosamente"})
}

func (ctrl *SalidasController) GetSalidasByCendisHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		Id_cendis int32 `json:"id_cendis"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if request.Id_cendis == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID de cendis no proporcionado"})
		return
	}

	salidas, err := ctrl.GetSalidasByCendisUseCase.Execute(request.Id_cendis)
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

func (ctrl *SalidasController) DeleteSalidaHandler(w http.ResponseWriter, r *http.Request) {
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

func (ctrl *SalidasController) GetSalidasPendientesHandler(w http.ResponseWriter, r *http.Request) {

}

func (ctrl *SalidasController) AddToSalidaHandler(w http.ResponseWriter, r *http.Request) {

}

func (ctrl *SalidasController) CerrarSalidaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut && r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Estructura para recibir el ID del body
	var requestBody struct {
		IdSalida int32 `json:"id_salida"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Datos inválidos"})
		return
	}

	// Validar que el ID sea válido
	if requestBody.IdSalida <= 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "ID de salida inválido"})
		return
	}

	// Ejecutar el caso de uso
	if err := ctrl.CerrarSalidaUseCase.Execute(requestBody.IdSalida); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Salida cerrada exitosamente"})
}
