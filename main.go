package main

import (
    "log"
    "net/http"
    "os"

    "DIMISA/src/core/mysql"
    "github.com/joho/godotenv"
)

func main() {
    // Cargar variables de entorno
    err := godotenv.Load()
    if err != nil {
        log.Println("⚠️ No se encontró archivo .env, usando variables de entorno del sistema")
    }

    // Inicializar conexión a MySQL
    db := mysql.InitMySQL()

    // Registrar todas las rutas pasando la conexión
    mysql.RegisterRoutes(db)

    // Leer puerto del .env o usar default 8080
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
