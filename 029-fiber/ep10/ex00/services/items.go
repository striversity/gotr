package services

import (
	"fmt"
	"log/slog"
)

type ItemRepository struct {
	db map[string]Item
}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{
		db: make(map[string]Item),
	}
}

// Item represents an item in the db
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Desc  string  `json:"desc"`
}

func (ir *ItemRepository) FindById(id string) (Item, error) {
	item := ir.db[id]

	if item.ID == "" {
		return Item{}, fmt.Errorf("item not found")
	}

	return item, nil
}

func (ir *ItemRepository) Create(name string, price float64, desc string) (Item, error) {
	if name == "" || price <= 0 {
		return Item{}, fmt.Errorf("item requires valid name and price")
	}

	item := Item{
		Name:  name,
		Price: price,
		Desc:  desc,
	}

	item.ID = fmt.Sprintf("%v", len(ir.db)+1)
	ir.db[item.ID] = item
	slog.Info("item created", "id", item.ID)
	return item, nil
}

func (ir *ItemRepository) GetItems() []Item {

	var items = []Item{}
	for _, item := range ir.db {
		items = append(items, item)
	}

	return items
}
