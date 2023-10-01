package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Item represents an item
type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Desc  string `json:"desc"`
}

var db = make(map[string]Item)

// UpdateItemById updates an item if found
func UpdateItemById(c *fiber.Ctx) error {
  id:= c.Params("id")
	if id == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	type ItemCreateRequest struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
		Desc  string `json:"desc"`
	}
	var reqItem ItemCreateRequest
	if err := c.BodyParser(&reqItem); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	item:= db[id]
	if item.ID == "" {
		return c.SendStatus(http.StatusNotFound)
	}

	item.Name = reqItem.Name
	item.Price = reqItem.Price
	item.Desc = reqItem.Desc
	db[id] = item // update db

	return c.JSON(item)
}

// SearchItems returns all items that match the search query
func SearchItems(c *fiber.Ctx) error {
	q := c.Query("q")
	slog.Info("request for search", "q", q)
	if q == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	var items = []Item{}
	for _, item := range db {
		if strings.Contains(item.Name, q) || strings.Contains(item.Desc, q) {
			items = append(items, item)
		}
	}

	return c.JSON(items)
}

// DeleteItemById returns an item if found
func DeleteItemById(c *fiber.Ctx) error {
	id := c.Params("id")
	slog.Info("request to delete item", "id", id)
	item := db[id]

	if item.ID == "" {
		return c.SendStatus(http.StatusNotFound)
	}

	delete(db, id)

	return c.SendStatus(http.StatusOK)
}

// GetItemById returns an item if found
func GetItemById(c *fiber.Ctx) error {
	id := c.Params("id")
	slog.Info("request for item", "id", id)
	item := db[id]

	if item.ID == "" {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(item)
}

// CreateItems creates an item
func CreateItems(c *fiber.Ctx) error {
	type ItemCreateRequest struct {
		Name  string `json:"name"`
		Price int    `json:"price"`
		Desc  string `json:"desc"`
	}
	var reqItem ItemCreateRequest
	if err := c.BodyParser(&reqItem); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	item := Item{
		Name:  reqItem.Name,
		Price: reqItem.Price,
		Desc:  reqItem.Desc,
	}

	item.ID = fmt.Sprintf("%v", len(db)+1)
	db[item.ID] = item
	slog.Info("item created", "id", item.ID)
	return c.JSON(item)
}

// GetItems returns all items
func GetItems(c *fiber.Ctx) error {
	slog.Info("request for all items")

	var items = []Item{}
	for _, item := range db {
		items = append(items, item)
	}

	return c.JSON(items)
}
