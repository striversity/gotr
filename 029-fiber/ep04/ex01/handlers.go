package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func getAllBooks(c *fiber.Ctx) error {
	slog.Info("request to get all books")
	return nil
}
