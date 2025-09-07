package main

import (
	"log"
	"net/http"
	"os"

	"DIMISA/src/core/mysql"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No se encontró archivo .env, usando variables de entorno del sistema")
	}

	db := mysql.InitMySQL()

	mysql.RegisterRoutes(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor corriendo en http://localhost:%s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("❌ Error al iniciar el servidor: %v", err)
	}
}
