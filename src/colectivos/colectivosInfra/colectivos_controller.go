package colectivosInfra

import (
	"encoding/json"
	"net/http"

	"DIMISA/src/colectivos/colectivosApp"
	"DIMISA/src/colectivos/colectivosDomain/colectivoEntity"
)

type ColectivosController struct {
	CreateColectivoUC       *colectivosApp.CreateColectivo
	GetColectivosByCendisUC *colectivosApp.GetColectivosByCendis
	GetPendingByCendisUC    *colectivosApp.GetPendingColectivosByCendis
	GetUpdatablesByCendisUC *colectivosApp.GetUpdatableColectivosByCendis
	AddToColectivoUC        *colectivosApp.AddToColectivo
	CloseColectivoUC        *colectivosApp.CloseColectivo
}

func NewColectivosController(
	createUC *colectivosApp.CreateColectivo,
	getByCendisUC *colectivosApp.GetColectivosByCendis,
	getPendingByCendisUC *colectivosApp.GetPendingColectivosByCendis,
	getUpdatablesByCendisUC *colectivosApp.GetUpdatableColectivosByCendis,
	addToColectivoUC *colectivosApp.AddToColectivo,
	closeColectivoUC *colectivosApp.CloseColectivo,
) *ColectivosController {
	return &ColectivosController{
		CreateColectivoUC:       createUC,
		GetColectivosByCendisUC: getByCendisUC,
		GetPendingByCendisUC:    getPendingByCendisUC,
		GetUpdatablesByCendisUC: getUpdatablesByCendisUC,
		AddToColectivoUC:        addToColectivoUC,
		CloseColectivoUC:        closeColectivoUC,
	}
}

func (cc *ColectivosController) CreateColectivoHandler(w http.ResponseWriter, r *http.Request) {
	var c colectivoEntity.ColectivoEntity
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := cc.CreateColectivoUC.Execute(&c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func (cc *ColectivosController) GetColectivosByCendisHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id_cendis int32 `json:"id_cendis"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	colectivos, err := cc.GetColectivosByCendisUC.Execute(body.Id_cendis)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(colectivos)
}

func (cc *ColectivosController) GetPendingColectivosByCendisHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id_cendis int32 `json:"id_cendis"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pendientes, err := cc.GetPendingByCendisUC.Execute(body.Id_cendis)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pendientes)
}

func (cc *ColectivosController) GetUpdatableColectivosByCendisHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id_cendis int32 `json:"id_cendis"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	editables, err := cc.GetUpdatablesByCendisUC.Execute(body.Id_cendis)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(editables)
}

func (cc *ColectivosController) AddToColectivoHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id_cendis     int32                                     `json:"id_cendis"`
		TipoColectivo int32                                     `json:"tipo_colectivo"`
		Detalles      []*colectivoEntity.ColectivoDetalleEntity `json:"detalles"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Error al decodificar el request: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validaciones
	if body.TipoColectivo <= 0 {
		http.Error(w, "tipo_colectivo debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	if len(body.Detalles) == 0 {
		http.Error(w, "detalles no puede estar vacío", http.StatusBadRequest)
		return
	}

	// Ejecutar caso de uso
	if err := cc.AddToColectivoUC.Execute(body.Id_cendis, body.TipoColectivo, body.Detalles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Detalles agregados correctamente al colectivo",
	})
}

func (cc *ColectivosController) CloseColectivoHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID int32 `json:"id_colectivo"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := cc.CloseColectivoUC.Execute(req.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Colectivo cerrado exitosamente"})
}
