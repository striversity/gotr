package handlers

import (
	"awesome/ex00/services"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ItemApi struct {
	repo *services.ItemRepository
}

func NewItemApi(ir *services.ItemRepository) *ItemApi {
	return &ItemApi{
		repo: ir,
	}
}

// GetItemById returns an item if found
func (ia *ItemApi) GetItemById(c *fiber.Ctx) error {
	id := c.Params("id")
	slog.Info("request for item", "id", id)

	if item, err := ia.repo.FindById(id); err != nil {
		slog.Error("error finding item", "path", c.Path(), "error", err)
		return c.SendStatus(http.StatusNotFound)
	} else {
		return c.JSON(item)
	}
}

// CreateItems creates an item
func (ia *ItemApi) CreateItems(c *fiber.Ctx) error {
	type ItemCreateRequest struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
		Desc  string  `json:"desc"`
	}
	var reqItem ItemCreateRequest
	if err := c.BodyParser(&reqItem); err != nil {
		slog.Error("error parsing request", "path", c.Path(), "error", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	if item, err := ia.repo.Create(reqItem.Name, reqItem.Price, reqItem.Desc); err != nil {
		slog.Error("error creating item", "path", c.Path(), "error", err)
		return c.SendStatus(http.StatusInternalServerError)
	} else {
		return c.JSON(item)
	}
}

// GetItems returns all items
func (ia *ItemApi) GetItems(c *fiber.Ctx) error {
	slog.Info("request for all items")

	items := ia.repo.GetItems()
	slog.Info("found items", "path", c.Path(), "count", len(items))

	return c.JSON(items)
}
