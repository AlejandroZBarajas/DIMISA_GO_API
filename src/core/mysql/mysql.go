package mysql

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

func InitMySQL() *sql.DB {
	
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("❌ Error conectando a MySQL: %v", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("❌ MySQL no responde: %v", err)
    }

    log.Println("✅ Conectado a MySQL")
    return db
}
