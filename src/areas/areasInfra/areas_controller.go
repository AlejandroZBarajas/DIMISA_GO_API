package areasInfra

import (
	"DIMISA/src/areas/areasApp"
	"DIMISA/src/areas/areasDomain/areaEntity"
	"encoding/json"
	"net/http"
)

type AreasController struct {
	CreateUC  *areasApp.CreateAreaUseCase
	UpdateUC  *areasApp.UpdateAreaUseCase
	GetAllUC  *areasApp.GetAllAreasUseCase
	GetByIDUC *areasApp.GetAreaByIDUseCase
	DeleteUC  *areasApp.DeleteAreaUseCase
}

func NewAreasController(
	createUC *areasApp.CreateAreaUseCase,
	updateUC *areasApp.UpdateAreaUseCase,
	getAllUC *areasApp.GetAllAreasUseCase,
	getByIDUC *areasApp.GetAreaByIDUseCase,
	deleteUC *areasApp.DeleteAreaUseCase,
) *AreasController {
	return &AreasController{
		CreateUC:  createUC,
		UpdateUC:  updateUC,
		GetAllUC:  getAllUC,
		GetByIDUC: getByIDUC,
		DeleteUC:  deleteUC,
	}
}

func (c *AreasController) CreateAreaHandler(w http.ResponseWriter, r *http.Request) {
	var area areaEntity.AreaEntity
	if err := json.NewDecoder(r.Body).Decode(&area); err != nil {
		http.Error(w, "Error en el cuerpo de la petición: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.CreateUC.Execute(&area); err != nil {
		http.Error(w, "Error al crear el área: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(area)
}

func (c *AreasController) UpdateAreaHandler(w http.ResponseWriter, r *http.Request) {
	var area areaEntity.AreaEntity
	if err := json.NewDecoder(r.Body).Decode(&area); err != nil {
		http.Error(w, "Error al parsear el cuerpo", http.StatusBadRequest)
		return
	}

	if err := c.UpdateUC.Execute(&area); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(area)
}

func (c *AreasController) GetAllAreasHandler(w http.ResponseWriter, r *http.Request) {
	areas, err := c.GetAllUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(areas)
}

func (c *AreasController) GetAreaByIDHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id_area int32 `json:"id_area"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al parsear el cuerpo", http.StatusBadRequest)
		return
	}

	area, err := c.GetByIDUC.Execute(req.Id_area)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if area == nil {
		http.Error(w, "Área no encontrada", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(area)
}

func (c *AreasController) DeleteAreaHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Id_area int32 `json:"id_area"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al parsear el cuerpo", http.StatusBadRequest)
		return
	}

	if err := c.DeleteUC.Execute(req.Id_area); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Área eliminada correctamente"}`))
}
