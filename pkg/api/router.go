package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/go-cleanar/internal/application/services"
	"github.com/hectorcoellomx/go-cleanar/internal/application/usecases"

	"github.com/hectorcoellomx/go-cleanar/internal/infrastructure/handlers"
	"github.com/hectorcoellomx/go-cleanar/internal/infrastructure/persistence/repositories"
	"github.com/hectorcoellomx/go-cleanar/pkg/api/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository) /*userService := &user.UserService{UserRepository: userRepository}*/
	getUserUsecase := usecases.NewGetUsers(*userService)

	userHandler := handlers.NewUserHandler(*getUserUsecase)
	log.Print(userHandler)

	api := app.Group("/api")

	api.Get("/users", userHandler.GetUsers)

	// Rutas de usuario
	//userRoutes := api.Group("/users")
	//userRoutes.Post("/", userHandler.CreateUser)
	//userRoutes.Get("/:id", userHandler.GetUserByID)

	api.Use(middleware.JWTMiddleware())
}
