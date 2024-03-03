package main

import (
	"awesome/ex02/server/endpoints"
	"awesome/ex02/server/storage"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	addr = ":8081"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		Immutable:     true,
		CaseSensitive: false,
	})

	api := app.Group("/api", logger.New())

	// configure store and endpoint for Items
	itemsStore := storage.NewItemsStore()
	itemsEndpoint := endpoints.NewItemsResource(itemsStore)

	// register items endpoints
	api.Post("/items", itemsEndpoint.CreateItem)
	api.Get("/items/:id", itemsEndpoint.GetItemById)

	err := app.Listen(addr)
	slog.Error("error starting server", "address", addr, "err", err)
}
