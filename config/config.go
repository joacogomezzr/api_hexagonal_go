package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv carga las variables del archivo .env
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠ No se pudo cargar el archivo .env, usando variables de entorno del sistema")
	} else {
		log.Println("✅ Archivo .env cargado correctamente")
	}
}

// GetEnv obtiene una variable de entorno o devuelve un valor por defecto
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
