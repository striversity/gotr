package main

import (
	"awesome/ex01/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
	}

	app := fiber.New(appConfig)
	app.Get("/items", handlers.GetItems)
	app.Post("/items", handlers.CreateItems)
	app.Get("/items/:id<int>", handlers.GetItemById)
	app.Delete("/items/:id<int>", handlers.DeleteItemById)
	app.Get("/items/search", handlers.SearchItems) // :3000/items/search?q=query_string
	app.Put("/items/:id<int>", handlers.UpdateItemById)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
