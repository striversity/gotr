package main

import (
	"awesome/ex01/handlers"
	"awesome/ex01/middlewares"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
	}

	app := fiber.New(appConfig)
	app.Use(middlewares.RequestID, middlewares.Logging)
	app.Get("/items", middlewares.VerifyLogin, handlers.GetItems)
	app.Post("/login", handlers.DoLogin)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
