package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/go-cleanar/internal/application/user"
	"github.com/hectorcoellomx/go-cleanar/internal/infrastructure/persistence/repositories"
	"github.com/hectorcoellomx/go-cleanar/pkg/api/handlers"
	"github.com/hectorcoellomx/go-cleanar/pkg/api/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	userRepository := repositories.NewUserRepository(db)

	userService := &user.UserService{UserRepository: userRepository}
	//userService := &user.UserService.NewUserService(userRepository)

	userUseCase := &user.userUseCase.NewUserUseCase(userService)

	/*userHandler := &handlers.UserHandler{ UserUseCase: userUseCase }*/
	userHandler := &handlers.UserHandler.NewUserHandler(userUseCase, userUseCase)

	api := app.Group("/api")

	// Rutas de usuario
	userRoutes := api.Group("/users")
	userRoutes.Post("/", userHandler.CreateUser)
	userRoutes.Get("/:id", userHandler.GetUserByID)

	// Otras rutas de entidad, por ejemplo, productos
	// ...

	// Middleware de autenticación JWT
	api.Use(middleware.JWTMiddleware())
}
