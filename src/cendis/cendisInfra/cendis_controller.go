package cendisInfra

import (
	"DIMISA/src/cendis/cendisApp"
	cendisEntity "DIMISA/src/cendis/cendisDomain/entity"
	"encoding/json"
	"net/http"
	"strconv"
)

type CendisController struct {
	CreateUC *cendisApp.CreateCendisUseCase
	UpdateUC *cendisApp.UpdateCendisUseCase
	GetAllUC *cendisApp.GetAllCendisUseCase
	DeleteUC *cendisApp.DeleteCendisUseCase
}

func NewCendisController(
	createUC *cendisApp.CreateCendisUseCase,
	updateUC *cendisApp.UpdateCendisUseCase,
	getAllUC *cendisApp.GetAllCendisUseCase,
	deleteUC *cendisApp.DeleteCendisUseCase,
) *CendisController {
	return &CendisController{
		CreateUC: createUC,
		UpdateUC: updateUC,
		GetAllUC: getAllUC,
		DeleteUC: deleteUC,
	}
}

type CreateCendisRequest struct {
	CendisNombre string  `json:"cendis_nombre"`
	Areas        []int32 `json:"areas"`
}

func (c *CendisController) CreateCendisHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateCendisRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al decodificar request", http.StatusBadRequest)
		return
	}

	if len(req.Areas) == 0 {
		http.Error(w, "Debe seleccionar al menos un área", http.StatusBadRequest)
		return
	}

	cendis := &cendisEntity.CendisEntity{
		Cendis_nombre: req.CendisNombre,
	}

	err := c.CreateUC.Repo.(*CendisRepository).CreateCendis(cendis, req.Areas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cendis)
}

func (c *CendisController) UpdateCendisHandler(w http.ResponseWriter, r *http.Request) {
	// Estructura para recibir el body completo
	var payload struct {
		Id_cendis     int32   `json:"id_cendis"`
		Cendis_nombre string  `json:"cendis_nombre"`
		Areas         []int32 `json:"areas"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Error al decodificar request", http.StatusBadRequest)
		return
	}

	// Validar que al menos haya un área
	if len(payload.Areas) == 0 {
		http.Error(w, "Debe asociar al menos un área al cendis", http.StatusBadRequest)
		return
	}

	// Construir la entidad
	cendis := cendisEntity.CendisEntity{
		Id_cendis:     payload.Id_cendis,
		Cendis_nombre: payload.Cendis_nombre,
	}

	// Ejecutar caso de uso con entidad y áreas
	if err := c.UpdateUC.Execute(&cendis, payload.Areas); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder con el objeto actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Cendis cendisEntity.CendisEntity `json:"cendis"`
		Areas  []int32                   `json:"areas"`
	}{
		Cendis: cendis,
		Areas:  payload.Areas,
	})
}

func (c *CendisController) GetAllCendisHandler(w http.ResponseWriter, r *http.Request) {
	cendisList, err := c.GetAllUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cendisList)
}

func (c *CendisController) DeleteCendisHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := c.DeleteUC.Execute(int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cendis eliminado"})
}
