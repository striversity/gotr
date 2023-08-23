package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func getAllBooksLC(c *fiber.Ctx) error {
	slog.Info("request to get all books - lowercase")
	return nil
}

func getAllBooksUC(c *fiber.Ctx) error {
	slog.Info("request to get all books - uppercase")
	return nil
}

func getBookById(c *fiber.Ctx) error {
	bookId := c.Params("id")
	slog.Info("request to get a book by id", "bookId", bookId)
	return nil
}
