package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/go-cleanar/internal/application/services"
	"github.com/hectorcoellomx/go-cleanar/internal/application/usecases/auth"
	"github.com/hectorcoellomx/go-cleanar/internal/application/usecases/user"

	"github.com/hectorcoellomx/go-cleanar/internal/infrastructure/handlers"
	"github.com/hectorcoellomx/go-cleanar/internal/infrastructure/repositories"
	"github.com/hectorcoellomx/go-cleanar/pkg/api/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	jwtmid := middleware.JWTMiddleware

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	getUserUsecase := user.NewGetUsers(*userService)
	userHandler := handlers.NewUserHandler(*getUserUsecase)

	loginUsecase := auth.NewLogin(*userService)
	loginHandler := handlers.NewAuthHandler(*loginUsecase)

	api := app.Group("/api")

	api.Get("/login", loginHandler.Login)
	api.Get("/refresh-token", loginHandler.RefreshToken)

	api.Get("/users", jwtmid, userHandler.GetUsers)

	//api.Use(middleware.JWTMiddlewareHandler())
}
