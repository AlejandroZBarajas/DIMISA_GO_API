package mysql

import (
	"DIMISA/src/areas/areasApp"
	"DIMISA/src/areas/areasInfra"
	"DIMISA/src/camas/camasApp"
	"DIMISA/src/camas/camasInfra"
	"DIMISA/src/users/userApp"
	"DIMISA/src/users/userInfra"
	"database/sql"
	"log"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {

	userRepo := &userInfra.UserRepository{DB: db}

	createUserUC := &userApp.CreateUserUseCase{Repo: userRepo}
	updateUserUC := &userApp.UpdateUserUseCase{Repo: userRepo}
	deleteUserUC := &userApp.DeleteUserUseCase{Repo: userRepo}
	getAllUserUC := &userApp.GetAllUsersUseCase{Repo: userRepo}
	getByRolUserUC := &userApp.GetUsersByRolUseCase{Repo: userRepo}
	getByIdUserUC := &userApp.GetUserByIDUseCase{Repo: userRepo}
	getByAreaUserUC := &userApp.GetUsersByAreaUseCase{Repo: userRepo}
	getByCendisUserUC := &userApp.GetUsersByCendisUseCase{Repo: userRepo}

	userController := userInfra.NewUserController(
		createUserUC,
		updateUserUC,
		deleteUserUC,
		getAllUserUC,
		getByRolUserUC,
		getByIdUserUC,
		getByAreaUserUC,
		getByCendisUserUC,
	)

	http.HandleFunc("/users/create", userController.CreateUserHandler)          // POST
	http.HandleFunc("/users/update", userController.UpdateUserHandler)          // PUT
	http.HandleFunc("/users/delete", userController.DeleteUserHandler)          // DELETE
	http.HandleFunc("/users/all", userController.GetAllUsersHandler)            // POST (según tu regla)
	http.HandleFunc("/users/by-rol", userController.GetUsersByRolHandler)       // POST
	http.HandleFunc("/users/by-id", userController.GetUserByIDHandler)          // POST
	http.HandleFunc("/users/by-area", userController.GetUsersByAreaHandler)     // POST
	http.HandleFunc("/users/by-cendis", userController.GetUsersByCendisHandler) // POST

	log.Println("✅ Rutas de usuarios registradas")

	camaRepo := &camasInfra.CamaRepository{DB: db}

	createCamaUC := &camasApp.CreateCama{Repo: camaRepo}
	updateCamaUC := &camasApp.UpdateCama{Repo: camaRepo}
	deleteCamaUC := &camasApp.DeleteCama{Repo: camaRepo}
	getByAreaCamaUC := &camasApp.GetCamasByArea{Repo: camaRepo}
	enableCamaUC := &camasApp.EnableCama{Repo: camaRepo}
	disableCamaUC := &camasApp.DisableCama{Repo: camaRepo}

	camaController := camasInfra.NewCamaController(
		createCamaUC,
		updateCamaUC,
		deleteCamaUC,
		getByAreaCamaUC,
		enableCamaUC,
		disableCamaUC,
	)

	// Rutas
	http.HandleFunc("/camas/create", camaController.CreateCamaHandler)      // POST
	http.HandleFunc("/camas/update", camaController.UpdateCamaHandler)      // PUT
	http.HandleFunc("/camas/delete", camaController.DeleteCamaHandler)      // DELETE
	http.HandleFunc("/camas/by-area", camaController.GetCamasByAreaHandler) // POST
	http.HandleFunc("/camas/enable", camaController.EnableCamaHandler)      // PUT
	http.HandleFunc("/camas/disable", camaController.DisableCamaHandler)    // PUT

	log.Println("✅ Rutas de camas registradas")

	areaRepo := &areasInfra.AreasRepository{DB: db}

	createAreaUC := &areasApp.CreateAreaUseCase{
		Repo:     areaRepo,
		CamaRepo: camaRepo,
	}
	updateAreaUC := &areasApp.UpdateAreaUseCase{Repo: areaRepo}
	getAllAreaUC := &areasApp.GetAllAreasUseCase{Repo: areaRepo}
	getByIDAreaUC := &areasApp.GetAreaByIDUseCase{Repo: areaRepo}
	deleteAreaUC := &areasApp.DeleteAreaUseCase{Repo: areaRepo}

	areaController := areasInfra.NewAreasController(
		createAreaUC,
		updateAreaUC,
		getAllAreaUC,
		getByIDAreaUC,
		deleteAreaUC,
	)

	http.HandleFunc("/areas/create", areaController.CreateAreaHandler) // POST
	http.HandleFunc("/areas/update", areaController.UpdateAreaHandler) // PUT
	http.HandleFunc("/areas/delete", areaController.DeleteAreaHandler) // DELETE
	http.HandleFunc("/areas/all", areaController.GetAllAreasHandler)   // POST
	http.HandleFunc("/areas/by-id", areaController.GetAreaByIDHandler) // POST

	log.Println("✅ Rutas de áreas registradas")
}
