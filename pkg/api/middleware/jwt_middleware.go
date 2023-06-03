package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/go-cleanar/internal/infrastructure/jwt"
)

func JWTMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	validate := jwt.ValidateToken(c, authHeader)

	if validate["success"] == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": validate["message"], "error_code": validate["errorno"]})
	}

	return c.Next()
}

func JWTMiddlewareHandler() fiber.Handler {
	return JWTMiddleware
}
