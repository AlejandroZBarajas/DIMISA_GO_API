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
}

func NewColectivosController(
	createUC *colectivosApp.CreateColectivo,
	getByCendisUC *colectivosApp.GetColectivosByCendis,
	getPendingByCendisUC *colectivosApp.GetPendingColectivosByCendis,
	getUpdatablesByCendisUC *colectivosApp.GetUpdatableColectivosByCendis,

) *ColectivosController {
	return &ColectivosController{
		CreateColectivoUC:       createUC,
		GetColectivosByCendisUC: getByCendisUC,
		GetPendingByCendisUC:    getPendingByCendisUC,
		GetUpdatablesByCendisUC: getUpdatablesByCendisUC,
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
