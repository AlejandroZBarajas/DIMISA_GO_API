package mysql

import (
		"database/sql"
	"log"
	"net/http"
	"DIMISA/src/users/userInfra"
	"DIMISA/src/users/userApp"
)

func RegisterRoutes(db *sql.DB) {
	// Repositorio
	userRepo := &userInfra.UserRepository{DB: db}

	// UseCase
	createUseCase := &userApp.CreateUserUseCase{Repo: userRepo}

	// Controlador
	userController := userInfra.NewUserController(createUseCase)

	// Rutas
	http.HandleFunc("/users/create", userController.CreateUserHandler)


	log.Println("✅ Rutas de usuarios registradas")
}
