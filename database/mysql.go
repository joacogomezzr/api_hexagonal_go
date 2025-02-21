package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)


var DB *sql.DB


func InitDB() error {
	
	requiredVars := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Error: La variable de entorno %s no está configurada", v)
		}
	}


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)


	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error conectando a la base de datos:", err)
		return err
	}

	
	if err := db.Ping(); err != nil {
		log.Println("Error verificando la conexión a la base de datos:", err)
		return err
	}

	DB = db
	log.Println("✅ Conexión a MySQL exitosa")
	return nil
}
