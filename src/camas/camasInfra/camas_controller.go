package camasInfra

import (
	"DIMISA/src/camas/camasApp"
	"DIMISA/src/camas/camasDomain/camaEntity"
	"encoding/json"
	"fmt"
	"net/http"
)

type CamaController struct {
	CreateUC        *camasApp.CreateCama
	UpdateUC        *camasApp.UpdateCama
	DeleteUC        *camasApp.DeleteCama
	GetByAreaUC     *camasApp.GetCamasByArea
	EnableUC        *camasApp.EnableCama
	DisableUC       *camasApp.DisableCama
	CreateRangeUC   *camasApp.CreateCamasRange
	GetFreeByAreaUC *camasApp.GetFreeCamasByArea
	SetFreeCamaUC   *camasApp.SetFreeCama
}

func NewCamaController(
	create *camasApp.CreateCama,
	update *camasApp.UpdateCama,
	deleteCama *camasApp.DeleteCama,
	getByArea *camasApp.GetCamasByArea,
	enable *camasApp.EnableCama,
	disable *camasApp.DisableCama,
	createRange *camasApp.CreateCamasRange,
	getFreeByArea *camasApp.GetFreeCamasByArea,
	setFreeCama *camasApp.SetFreeCama,

) *CamaController {
	return &CamaController{
		CreateUC:        create,
		UpdateUC:        update,
		DeleteUC:        deleteCama,
		GetByAreaUC:     getByArea,
		EnableUC:        enable,
		DisableUC:       disable,
		CreateRangeUC:   createRange,
		GetFreeByAreaUC: getFreeByArea,
		SetFreeCamaUC:   setFreeCama,
	}
}

func (c *CamaController) CreateCamasRangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Id_area int32 `json:"id_area"`
		Cama_1  int32 `json:"cama_1"`
		Cama_n  int32 `json:"cama_n"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	usecase := &camasApp.CreateCamasRange{Repo: c.CreateUC.Repo} // asumimos que Repo es accesible
	if err := usecase.Execute(input.Id_area, input.Cama_1, input.Cama_n); err != nil {
		http.Error(w, fmt.Sprintf("Error al crear camas: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Camas creadas correctamente"))
}

func (c *CamaController) CreateCamaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input camaEntity.CamaEntity
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	err := c.CreateUC.Execute(&input)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al crear cama: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Cama creada exitosamente")))
}

func (c *CamaController) UpdateCamaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("entra al servicio para editar la cama en el back")
	var input camaEntity.CamaEntity
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	err := c.UpdateUC.Execute(&input)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar cama: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cama actualizada exitosamente"})
}

func (c *CamaController) DeleteCamaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct{ Id_cama int32 }
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	err := c.DeleteUC.Execute(input.Id_cama)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al eliminar cama: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cama eliminada exitosamente"})
}

func (c *CamaController) EnableCamaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct{ Id_cama int32 }
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	err := c.EnableUC.Execute(input.Id_cama)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al habilitar cama: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cama habilitada exitosamente"})
}

func (c *CamaController) DisableCamaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct{ Id_cama int32 }
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	err := c.DisableUC.Execute(input.Id_cama)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al deshabilitar cama: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Cama deshabilitada exitosamente"})
}

func (c *CamaController) GetCamasByAreaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct{ Id_area int32 }
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	camas, err := c.GetByAreaUC.Execute(input.Id_area)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener camas: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(camas)
}

func (c *CamaController) GetFreeCamasByAreaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		IdArea int32 `json:"id_area"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	camas, err := c.GetFreeByAreaUC.Execute(body.IdArea)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener camas: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(camas)
}

func (c *CamaController) SetFreeCamaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var body struct {
		IdCama int32 `json:"id_cama"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err := c.SetFreeCamaUC.Execute(body.IdCama)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error al desocupar cama: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

}
