package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func Api(c *fiber.Ctx) error {
	slog.Info("middleware", "name", "/api")
	return c.Next()
}

func ApiV1(c *fiber.Ctx) error {
	slog.Info("middleware", "name", "/v1")
	return c.Next()
}

func ApiV2(c *fiber.Ctx) error {
	slog.Info("middleware", "name", "/v2")
	return c.Next()
}
