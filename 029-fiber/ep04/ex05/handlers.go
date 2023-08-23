package main

import (
	"log/slog"
	"strings"

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

func getAllAuthorsOrAuthorById(c *fiber.Ctx) error {
	authorId := c.Params("id")

	if authorId == "" {
		slog.Info("request all authors")
		return nil
	}

	slog.Info("request to get an author by id", "authorId", authorId)
	return nil
}

func getItems(c *fiber.Ctx) error {
	itemPath := c.Params("*")

	if itemPath == "" {
		slog.Info("request all items")
		return nil
	}

	slog.Info("request to get an item with sub-paths:")
	subPaths := strings.Split(itemPath, "/")
	for _, subPath := range subPaths {
		slog.Info("\titem subPath", "subPath", subPath)
	}

	return nil
}
