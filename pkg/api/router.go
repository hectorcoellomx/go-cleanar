package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/go-cleanar/internal/application/services"
	"github.com/hectorcoellomx/go-cleanar/internal/application/usecases"

	"github.com/hectorcoellomx/go-cleanar/internal/infrastructure/handlers"
	"github.com/hectorcoellomx/go-cleanar/internal/infrastructure/repositories"
	"github.com/hectorcoellomx/go-cleanar/pkg/api/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	getUserUsecase := usecases.NewGetUsers(*userService)
	userHandler := handlers.NewUserHandler(*getUserUsecase)

	api := app.Group("/api")
	api.Get("/users", userHandler.GetUsers)

	api.Use(middleware.JWTMiddleware())
}
