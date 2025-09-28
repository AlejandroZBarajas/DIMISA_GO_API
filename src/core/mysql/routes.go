package mysql

import (
	"DIMISA/src/areas/areasApp"
	"DIMISA/src/areas/areasInfra"
	"DIMISA/src/camas/camasApp"
	"DIMISA/src/camas/camasInfra"
	"DIMISA/src/core/auth"
	"DIMISA/src/users/userApp"
	"DIMISA/src/users/userInfra"
	"database/sql"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func RegisterRoutes(db *sql.DB) {

	mux := http.NewServeMux()

	loginHandler := &auth.LoginHandler{DB: db}
	mux.Handle("/login", loginHandler)
	log.Println("✅ Ruta de login registrada")

	userRepo := &userInfra.UserRepository{DB: db}

	userController := userInfra.NewUserController(
		&userApp.CreateUserUseCase{Repo: userRepo},
		&userApp.UpdateUserUseCase{Repo: userRepo},
		&userApp.DeleteUserUseCase{Repo: userRepo},
		&userApp.GetAllUsersUseCase{Repo: userRepo},
		&userApp.GetUsersByRolUseCase{Repo: userRepo},
		&userApp.GetUserByIDUseCase{Repo: userRepo},
		&userApp.GetUsersByAreaUseCase{Repo: userRepo},
		&userApp.GetUsersByCendisUseCase{Repo: userRepo},
	)

	mux.HandleFunc("/users/create", userController.CreateUserHandler)          // POST
	mux.HandleFunc("/users/update", userController.UpdateUserHandler)          // PUT
	mux.HandleFunc("/users/delete", userController.DeleteUserHandler)          // DELETE
	mux.HandleFunc("/users/all", userController.GetAllUsersHandler)            // POST
	mux.HandleFunc("/users/by-rol", userController.GetUsersByRolHandler)       // POST
	mux.HandleFunc("/users/by-id", userController.GetUserByIDHandler)          // POST
	mux.HandleFunc("/users/by-area", userController.GetUsersByAreaHandler)     // POST
	mux.HandleFunc("/users/by-cendis", userController.GetUsersByCendisHandler) // POST

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
	mux.HandleFunc("/camas/create", camaController.CreateCamaHandler)      // POST
	mux.HandleFunc("/camas/update", camaController.UpdateCamaHandler)      // PUT
	mux.HandleFunc("/camas/delete", camaController.DeleteCamaHandler)      // DELETE
	mux.HandleFunc("/camas/by-area", camaController.GetCamasByAreaHandler) // POST
	mux.HandleFunc("/camas/enable", camaController.EnableCamaHandler)      // PUT
	mux.HandleFunc("/camas/disable", camaController.DisableCamaHandler)    // PUT

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
	getFreeAreasUC := &areasApp.GetFreeAreasUseCase{Repo: areaRepo}

	areaController := areasInfra.NewAreasController(
		createAreaUC,
		updateAreaUC,
		getAllAreaUC,
		getByIDAreaUC,
		deleteAreaUC,
		getFreeAreasUC,
	)

	mux.HandleFunc("/areas/create", areaController.CreateAreaHandler) // POST
	mux.HandleFunc("/areas/update", areaController.UpdateAreaHandler) // PUT
	mux.HandleFunc("/areas/delete", areaController.DeleteAreaHandler) // DELETE
	mux.HandleFunc("/areas", areaController.GetAllAreasHandler)       // POST
	mux.HandleFunc("/areas/by-id", areaController.GetAreaByIDHandler) // POST
	mux.HandleFunc("/areas/free", areaController.GetFreeAreasHandler)
	log.Println("✅ Rutas de áreas registradas")

	handlerWithCors := corsMiddleware(mux)

	log.Println("🚀 Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCors))
}
