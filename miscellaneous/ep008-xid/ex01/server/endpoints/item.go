package endpoints

import (
	"awesome/ex01/server/storage"
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

// Item represents a JSON serialized object between server and clients
type Item struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ItemsResource struct {
	store *storage.ItemsStore
}

func NewItemsResource(store *storage.ItemsStore) *ItemsResource {
	return &ItemsResource{
		store: store,
	}
}

// 1. Create an item - POST /api/items - with name, description and price fields
func (ep *ItemsResource) CreateItem(c *fiber.Ctx) error {
	item := &Item{}
	if err := c.BodyParser(item); err != nil {
		slog.Error("Error parsing request", "err", err)
		return err
	}

	id, err := ep.store.Save(item.Name, item.Description, item.Price)
	if err != nil {
		return err
	}

	return c.JSON(id, "text/plain")
}

// 2. Get an item - GET /api/items/:id
func (ep *ItemsResource) GetItemById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id", 0)

	slog.Info("retriving item by id", "id", id)
	pItem, err := ep.store.GetByID(int64(id))
	if err != nil {
		return err
	}

	slog.Info("got item", slog.Any("item", pItem))

	item := &Item{
		ID:          pItem.ID,
		Name:        pItem.Name,
		Description: pItem.Description,
		Price:       pItem.Price,
	}

	return c.JSON(item)
}

// 3. Update an item by id - PUT /api/item/{id} with name, description and price fields
// 4. Delete an item by id - DELETE /api/item/{id} with no body
