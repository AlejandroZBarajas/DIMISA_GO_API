package mysql

import (
	"DIMISA/src/users/userApp"
	"DIMISA/src/users/userInfra"
	"database/sql"
	"log"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {
	// Repositorio
	userRepo := &userInfra.UserRepository{DB: db}

	// UseCases
	createUC := &userApp.CreateUserUseCase{Repo: userRepo}
	updateUC := &userApp.UpdateUserUseCase{Repo: userRepo}
	deleteUC := &userApp.DeleteUserUseCase{Repo: userRepo}
	getAllUC := &userApp.GetAllUsersUseCase{Repo: userRepo}
	getByRolUC := &userApp.GetUsersByRolUseCase{Repo: userRepo}
	getByIdUC := &userApp.GetUserByIDUseCase{Repo: userRepo}
	getByAreaUC := &userApp.GetUsersByAreaUseCase{Repo: userRepo}
	getByCendisUC := &userApp.GetUsersByCendisUseCase{Repo: userRepo}

	// Controlador con todos los casos de uso
	userController := userInfra.NewUserController(
		createUC,
		updateUC,
		deleteUC,
		getAllUC,
		getByRolUC,
		getByIdUC,
		getByAreaUC,
		getByCendisUC,
	)

	// Rutas
	http.HandleFunc("/users/create", userController.CreateUserHandler)          // POST
	http.HandleFunc("/users/update", userController.UpdateUserHandler)          // PUT
	http.HandleFunc("/users/delete", userController.DeleteUserHandler)          // DELETE
	http.HandleFunc("/users/all", userController.GetAllUsersHandler)            // POST (según tu regla)
	http.HandleFunc("/users/by-rol", userController.GetUsersByRolHandler)       // POST
	http.HandleFunc("/users/by-id", userController.GetUserByIDHandler)          // POST
	http.HandleFunc("/users/by-area", userController.GetUsersByAreaHandler)     // POST
	http.HandleFunc("/users/by-cendis", userController.GetUsersByCendisHandler) // POST

	log.Println("✅ Rutas de usuarios registradas")
}
