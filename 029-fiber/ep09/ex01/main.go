package main

import (
	"awesome/ex01/handlers"
	"awesome/ex01/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
	}

	app := fiber.New(appConfig)
	api := app.Group("/api", middleware.Api)

	v1 := api.Group("/v1", middleware.ApiV1)
	v1.Get("/items", handlers.GetItems)
	v1.Post("/items", handlers.CreateItems)
	v1.Get("/items/:id<int>", handlers.GetItemById)

	v2 := api.Group("/v2", middleware.ApiV2)
	v2.Get("/items", handlers.GetItems)
	v2.Post("/items", handlers.CreateItems)
	v2.Get("/items/:id<int>", handlers.GetItemById)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
