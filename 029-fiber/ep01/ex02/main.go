package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

var items []Item = []Item{
	{1, "Item 1", 10.99},
	{2, "Item 2", 20.79},
	{3, "Item 3", 0.31},
}

// getAllItems returns all items
func getAllItems(c *fiber.Ctx) error {
	log.Infof("Received a %s request for items", c.Method())
	return c.JSON(items)
}

// getItem returns a single item
func getItem(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid ID")
	}

	for _, item := range items {
		if item.ID == id {
			return c.JSON(item)
		}
	}

	return c.Status(http.StatusNotFound).SendString("Item not found")
}

// createItem creates a new item
func createItem(c *fiber.Ctx) error {
	var item Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid JSON")
	}

	item.ID = len(items) + 1
	items = append(items, item)
	return c.JSON(item)
}

func main() {
	app := fiber.New()
	app.Get("/items", getAllItems)
	app.Get("/items/:id", getItem)
	app.Post("/items", createItem)

	log.Info("Starting server on port 3000")
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}

}
