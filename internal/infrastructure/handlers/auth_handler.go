package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/go-cleanar/internal/application/usecases/auth"
	"github.com/hectorcoellomx/go-cleanar/internal/infrastructure/jwt"
)

type AuthHandler struct {
	GetUsersUseCase auth.Login
}

func NewAuthHandler(getUsersUseCase auth.Login) *AuthHandler {
	return &AuthHandler{
		GetUsersUseCase: getUsersUseCase,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {

	token := jwt.GenerateToken(c, "1", "hector@gmail.com", "1", 1.0)
	token_refresh := jwt.GenerateToken(c, "1", "hector@gmail.com", "1", 168.0, "refresh")

	data := fiber.Map{
		"token":   token,
		"refresh": token_refresh,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Ok", "data": data})

}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	validate := jwt.ValidateToken(c, authHeader, "refresh")

	if validate["success"] == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": validate["message"], "error_code": validate["errorno"]})
	}

	claims := validate["claims"].(map[string]interface{})
	id := claims["sub"].(string)
	email := claims["email"].(string)
	role := claims["role"].(string)

	token := jwt.GenerateToken(c, id, email, role, 1.0)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Ok", "data": token})

}
