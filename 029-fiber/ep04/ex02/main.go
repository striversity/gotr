package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
		CaseSensitive: true,
	}

	app := fiber.New(appConfig)
	app.Get("/books", getAllBooksLC).Name("Get all books lowercase")
	app.Get("/Books", getAllBooksUC).Name("Get all books uppercase")

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
