package main

import (
	"log"
	"api-joaquin/config"
	"api-joaquin/database"
	"api-joaquin/pkg/middleware"
	"api-joaquin/internal/book/infrastructure/adapters"
	"os"
	
	"github.com/gofiber/fiber/v2"
	"fmt"

	// Importaciones para el recurso de libros
	bookControllers "api-joaquin/internal/book/controllers"
	bookInfrastructure "api-joaquin/internal/book/infrastructure"
	bookInterfaces "api-joaquin/internal/book/interface"
	bookRoutes "api-joaquin/internal/book/routes"

	// Importaciones para el recurso de administradores
	adminControllers "api-joaquin/internal/admin/controllers"
	adminInfrastructure "api-joaquin/internal/admin/infrastructure"
	adminInterfaces "api-joaquin/internal/admin/interface"
	adminRoutes "api-joaquin/internal/admin/routes"
)

func main() {
	// Cargar variables de entorno
	config.LoadEnv()

	// Verificar las variables de entorno
	requiredVars := []string{"DB_USER", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME"}
	for _, v := range requiredVars {
		if config.GetEnv(v, "") == "" {
			log.Fatalf("❌ Error: La variable de entorno %s no está configurada", v)
		}
	}

	// Inicializar la base de datos
	if err := database.InitDB(); err != nil {
		log.Fatal("Error al conectar con la base de datos")
	}

	// Inicializar Fiber
	app := fiber.New()
	// Configurar middleware CORS
	app.Use(middleware.SetupCORS())

	rabbitmqUser := os.Getenv("RABBITMQ_USER")
    rabbitmqPass := os.Getenv("RABBITMQ_PASS")
    rabbitmqHost := os.Getenv("RABBITMQ_HOST")
    rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPass, rabbitmqHost, rabbitmqPort)
	rabbitmqConn := adapters.NewRabbitMQAdapter(connStr)
	// Inicializar el repositorio y el controlador para libros
	bookRepo := bookInfrastructure.NewBookService(database.DB)
	bookController := bookControllers.NewBookController(bookRepo, rabbitmqConn)
	bookHandler := bookInterfaces.NewBookHandler(bookController)

	// Inicializar el repositorio y el controlador para administradores
	adminRepo := adminInfrastructure.NewAdminService(database.DB)
	adminController := adminControllers.NewAdminController(adminRepo)
	adminHandler := adminInterfaces.NewAdminHandler(adminController)

	// Configurar las rutas
	bookRoutes.SetupBookRoutes(app, bookHandler)
	adminRoutes.SetupAdminRoutes(app, adminHandler)

	// Obtener puerto desde configuración
	port := config.GetEnv("PORT", "8080")

	// Iniciar el servidor
	log.Printf("🚀 Servidor corriendo en el puerto %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
