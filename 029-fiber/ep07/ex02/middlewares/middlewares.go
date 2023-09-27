package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func VerifyLogin(c *fiber.Ctx) error {
	// ensure the user is logged in
	token := c.Request().Header.Peek("Authorization")
	if token == nil {
		return c.SendStatus(http.StatusUnauthorized)
	}

	jwt := string(token)
	if jwt != "Bearer 1234567890" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid token"})
	}

	slog.Info("got request", "method", c.Method(), "path", c.Path(), "token", token)
	return c.Next()
}
