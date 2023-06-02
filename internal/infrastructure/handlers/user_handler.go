package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/go-cleanar/internal/application/user"
)

type UserHandler struct {
	UserUseCase  user.UserUseCase
	UserUseCase2 user.UserUseCase
}

func NewUserHandler(userUseCase user.UserUseCase, userUseCase2 user.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase:  userUseCase,
		UserUseCase2: userUseCase2,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	// Obtener los datos del cuerpo de la solicitud
	type CreateUserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	req := new(CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		// Manejar el error de análisis del cuerpo de la solicitud
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	// Crear el usuario utilizando el caso de uso
	createdUser, err := h.UserUseCase.CreateUser(req.Name, req.Email)
	if err != nil {
		// Manejar el error de creación del usuario
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(createdUser)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	// Obtener el ID del parámetro de la ruta
	id := c.Params("id", "0")

	// Convertir el ID en un número entero
	userID, err := strconv.Atoi(id)
	if err != nil {
		// Manejar el error de conversión del ID
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	// Obtener el usuario por ID utilizando el caso de uso
	foundUser, err := h.UserUseCase.GetUserByID(uint(userID))
	if err != nil {
		// Manejar el error de búsqueda del usuario
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	if foundUser == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.JSON(foundUser)
}
