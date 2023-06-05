package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/go-cleanar/internal/application/usecases/user"
)

type UserHandler struct {
	GetUsersUseCase   user.GetUsers
	CreateUserUseCase user.CreateUser
}

func NewUserHandler(getUsersUseCase user.GetUsers, createUserUseCase user.CreateUser) *UserHandler {
	return &UserHandler{
		GetUsersUseCase:   getUsersUseCase,
		CreateUserUseCase: createUserUseCase,
	}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {

	users, err := h.GetUsersUseCase.GetUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 500})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Ok", "data": users})

}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	type CreateUserRequest struct {
		Id       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Status   int    `json:"status"`
	}

	req := new(CreateUserRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 400})
	}

	createdUser, err := h.CreateUserUseCase.CreateUser(req.Id, req.Username, req.Email, req.Password, req.Status)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 500})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Record created", "data": createdUser})
}

/*
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
*/
