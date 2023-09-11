package main

import (
	"log"
	"log/slog"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/gofiber/fiber/v2"
)

func getItemById(c *fiber.Ctx) error {
	item := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "Widget 1 for ACME",
	}

	return c.JSON(item)
}

func requestLogger(c *fiber.Ctx) error {
	// log request method, path, and param 'id'
	reqId := c.Request().Header.Peek("REQUEST-ID")
	slog.Info("got request", "method", c.Method(), "path", c.Path(), "id", c.Params("id"), "requestId", reqId)
	return c.Next()
}

func addRequestID(c *fiber.Ctx) error {
	// add a unique request ID to each request
	uid, _ := guid.NewV4()
	c.Request().Header.Add("REQUEST-ID", uid.String())
	return c.Next()
}

func main() {
	appConfig := fiber.Config{
		Immutable:         true,
		EnablePrintRoutes: true,
	}

	app := fiber.New(appConfig)
	app.Get("/items/:id", addRequestID, requestLogger, getItemById)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
