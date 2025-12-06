package entradasInfra

import (
	"DIMISA/src/entradas/entradasApp"
	"DIMISA/src/entradas/entradasDomain/entradaEntity"
	"encoding/json"
	"net/http"
)

type EntradasController struct {
	CapturarEntradaUC *entradasApp.CapturarEntradaUseCase
}

func NewEntradaController(capturar *entradasApp.CapturarEntradaUseCase) *EntradasController {
	return &EntradasController{
		CapturarEntradaUC: capturar,
	}
}

func (c *EntradasController) CapturarEntrada(w http.ResponseWriter, r *http.Request) {
	var entrada entradaEntity.EntradaRequest

	if err := json.NewDecoder(r.Body).Decode(&entrada); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if err := c.CapturarEntradaUC.Execute(&entrada); err != nil {
		http.Error(w, "Error al capturar entrada", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Entrada capturada correctamente"})
}
