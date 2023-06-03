package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hectorcoellomx/go-cleanar/internal/domain/entities"
	"github.com/hectorcoellomx/go-cleanar/pkg/api"
	"github.com/hectorcoellomx/go-cleanar/pkg/config"
)

func main() {

	app := fiber.New()

	cfg := config.LoadConfig()
	db, err := config.InitDatabase(cfg)

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entities.User{})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api.SetupRoutes(app, db)

	app.Listen(":8080")
}
