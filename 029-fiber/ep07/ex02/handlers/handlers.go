package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetItems is an authenticated endpoint that returns all items
func GetItems(c *fiber.Ctx) error {
	reqId:=c.Response().Header.Peek("X-Request-Id")
	slog.Info("get items request recieved", "path", c.Path(), "method", c.Method(), "reqId", string(reqId))
	return c.SendStatus(http.StatusOK)
}

// DoLogin is an unauthenticated endpoint that gets user info and tries to authenticate them
func DoLogin(c *fiber.Ctx) error {
	type loginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req loginReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if req.Username != "admin" || req.Password != "admin" {
		return c.Status(http.StatusUnauthorized).SendString("invalid credentials")
	}

	c.Response().Header.Add("Authorization", "Bearer 1234567890")
	slog.Info("login request recieved", "path", c.Path(), "method", c.Method())
	return c.SendStatus(http.StatusOK)
}
