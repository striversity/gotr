package main

import (
	"awesome/ex02/handlers"
	"awesome/ex02/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
	}

	app := fiber.New(appConfig)

	itemsRepo := services.NewItemRepository()
	itemApi := handlers.NewItemApi(itemsRepo)

	api := app.Group("/api")
	api.Get("/items", itemApi.GetItems)
	api.Post("/items", itemApi.CreateItems)
	api.Get("/items/:id<int>", itemApi.GetItemById)

	app.Static("/views", "./views")

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
