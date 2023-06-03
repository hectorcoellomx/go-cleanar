package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/go-cleanar/pkg/config"
)

func main() {
	app := fiber.New()

	cfg := config.LoadConfig()
	db, err := config.InitDatabase(cfg)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(db)

	// Configurar rutas y controladores
	//api.SetupRoutes(app, db)

	app.Listen(":8080")
}
