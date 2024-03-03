package storage

import (
	"fmt"
	"log/slog"
	"time"
)

// Item represents a presisted item value as a row in storage
type Item struct {
	ID int64
	// OwnerId int64
	Name        string
	Description string
	Price       float64
	CreatedOn   time.Time
	UpdatedOn   time.Time
}

type ItemsStore struct {
	items []Item
}

func NewItemsStore() *ItemsStore { return &ItemsStore{} }

// Save() saves an item object to the store and returns it's ID
func (s *ItemsStore) Save(name string, desc string, price float64) (int64, error) {
	// validate inputs

	item := Item{
		Name:        name,
		Description: desc,
		Price:       price,
		CreatedOn:   time.Now(),
		UpdatedOn:   time.Now(),
	}

	item.ID = int64(len(s.items) + 1)
	s.items = append(s.items, item)
	return item.ID, nil
}

// GetByID() saves an item object to the store and returns it's ID
func (s *ItemsStore) GetByID(id int64) (Item, error) {
	// validate inputs
	if id < 1 || id > int64(len(s.items)) {
		return Item{}, fmt.Errorf("invalid item id: %v", id)
	}

	item := s.items[id-1]
	slog.Info("retrieved item", "id", id, "item.name", item.Name)

	return item, nil
}
