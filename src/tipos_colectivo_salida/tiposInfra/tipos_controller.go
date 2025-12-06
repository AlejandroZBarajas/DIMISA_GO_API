package tiposInfra

import (
	"DIMISA/src/tipos_colectivo_salida/tiposApp"
	"encoding/json"
	"net/http"
)

type TiposController struct {
	GetTiposUC *tiposApp.GetTipos
}

func NewTiposController(getTiposUC *tiposApp.GetTipos) *TiposController {
	return &TiposController{
		GetTiposUC: getTiposUC,
	}
}

func (tc *TiposController) GetTiposHandler(w http.ResponseWriter, r *http.Request) {
	tipos, err := tc.GetTiposUC.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tipos)
}
