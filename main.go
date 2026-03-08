package main

import (
	"DIMISA/src/core/mysql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No se encontró archivo .env, usando variables de entorno del sistema")
	}

	db := mysql.InitMySQL()

	handler := mysql.RegisterRoutes(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor corriendo en http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
