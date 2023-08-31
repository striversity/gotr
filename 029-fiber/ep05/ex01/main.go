package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func getAllBooks(c *fiber.Ctx) error {
	// 2. Returning the result (responsing to the client)
	//    * return result or an error
	person := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		ID:   1,
		Name: "John",
		Age:  30,
	}

	return c.JSON(person)
}

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
	}

	app := fiber.New(appConfig)
	app.Get("/books", getAllBooks).Name("Get all books")

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
