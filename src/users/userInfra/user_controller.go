package userInfra

import (
	"DIMISA/src/users/userApp"
	"DIMISA/src/users/userDomain/usersEntities"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserController struct {
	CreateUseCase      *userApp.CreateUserUseCase
	UpdateUseCase      *userApp.UpdateUserUseCase
	GetAllUseCase      *userApp.GetAllUsersUseCase
	GetByRolUseCase    *userApp.GetUsersByRolUseCase
	GetByIDUseCase     *userApp.GetUserByIDUseCase
	DeleteUseCase      *userApp.DeleteUserUseCase
	GetByAreaUseCase   *userApp.GetUsersByAreaUseCase
	GetByCendisUseCase *userApp.GetUsersByCendisUseCase
}

func NewUserController(
	create *userApp.CreateUserUseCase,
	update *userApp.UpdateUserUseCase,
	deleteUC *userApp.DeleteUserUseCase,
	getAll *userApp.GetAllUsersUseCase,
	getByRol *userApp.GetUsersByRolUseCase,
	getById *userApp.GetUserByIDUseCase,
	getByArea *userApp.GetUsersByAreaUseCase,
	getByCendis *userApp.GetUsersByCendisUseCase,
) *UserController {
	return &UserController{
		CreateUseCase:      create,
		UpdateUseCase:      update,
		DeleteUseCase:      deleteUC,
		GetAllUseCase:      getAll,
		GetByRolUseCase:    getByRol,
		GetByIDUseCase:     getById,
		GetByAreaUseCase:   getByArea,
		GetByCendisUseCase: getByCendis,
	}
}

func (c *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Nombres   string `json:"nombres"`
		Apellido1 string `json:"apellido1"`
		Apellido2 string `json:"apellido2"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Id_rol    int32  `json:"id_rol"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	user := &usersEntities.UserEntity{
		Nombres:   input.Nombres,
		Apellido1: input.Apellido1,
		Apellido2: input.Apellido2,
		Username:  input.Username,
		Password:  input.Password,
		Id_rol:    input.Id_rol,
	}

	err := c.CreateUseCase.Execute(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al crear usuario: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario '%s' creado exitosamente", input.Username)))
}

func (c *UserController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Id_usuario int32  `json:"id_usuario"`
		Nombres    string `json:"nombres"`
		Apellido1  string `json:"apellido1"`
		Apellido2  string `json:"apellido2"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Id_rol     int32  `json:"id_rol"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	user := &usersEntities.UserEntity{
		Id_usuario: input.Id_usuario,
		Nombres:    input.Nombres,
		Apellido1:  input.Apellido1,
		Apellido2:  input.Apellido2,
		Username:   input.Username,
		Password:   input.Password,
		Id_rol:     input.Id_rol,
	}

	err := c.UpdateUseCase.Execute(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al actualizar usuario: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Usuario '%s' actualizado exitosamente", input.Username)))
}

func (c *UserController) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	users, err := c.GetAllUseCase.Execute()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener usuarios: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUsersByRolHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Id_rol int32 `json:"id_rol"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	users, err := c.GetByRolUseCase.Execute(input.Id_rol)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener usuarios por rol: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Id_usuario int32 `json:"id_usuario"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	user, err := c.GetByIDUseCase.Execute(input.Id_usuario)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener usuario: %v", err), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (c *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Id_usuario int32 `json:"id_usuario"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	err := c.DeleteUseCase.Execute(input.Id_usuario)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al eliminar usuario: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Usuario con ID %d eliminado exitosamente", input.Id_usuario)))
}

func (c *UserController) GetUsersByAreaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Id_area int32 `json:"id_area"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	users, err := c.GetByAreaUseCase.Execute(input.Id_area)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener usuarios por área: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (c *UserController) GetUsersByCendisHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Id_cendis int32 `json:"id_cendis"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	users, err := c.GetByCendisUseCase.Execute(input.Id_cendis)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener usuarios por cendis: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
