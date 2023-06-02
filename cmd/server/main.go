package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/go-cleanar/pkg/api"
	"github.com/hectorcoellomx/go-cleanar/pkg/config"
)

func main() {
	app := fiber.New()

	// Configurar dependencias
	cfg := config.LoadConfig()
	db := config.InitDatabase(cfg.DatabaseURL)

	// Configurar rutas y controladores
	api.SetupRoutes(app, db)

	// Iniciar el servidor
	log.Fatal(app.Listen(":8080"))
}
