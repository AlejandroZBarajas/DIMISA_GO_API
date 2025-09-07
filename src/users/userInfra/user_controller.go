package userInfra

import (
	"encoding/json"
	"net/http"
	"fmt"
	"DIMISA/src/users/userApp"
	"DIMISA/src/users/userDomain/usersEntities"
	"golang.org/x/crypto/bcrypt"

)

type UserController struct {
	CreateUseCase        *userApp.CreateUserUseCase
	
}

func NewUserController(
	create *userApp.CreateUserUseCase,
	
) *UserController {
	return &UserController{
		CreateUseCase:        create,
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
		Password  string `json:"password"` // recibir contraseña en texto plano
		Id_rol    int32  `json:"id_rol"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("Error al leer datos: %v", err), http.StatusBadRequest)
		return
	}

	// Hashear la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al hashear la contraseña: %v", err), http.StatusInternalServerError)
		return
	}

	// Crear la entidad
	user := &usersEntities.UserEntity{
		Nombres:   input.Nombres,
		Apellido1: input.Apellido1,
		Apellido2: input.Apellido2,
		Username:  input.Username,
		Password:  string(hashedPassword),
		Id_rol:    input.Id_rol,
	}

	// Ejecutar el use case
	err = c.CreateUseCase.Execute(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al crear usuario: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario '%s' creado exitosamente", input.Username)))
}
