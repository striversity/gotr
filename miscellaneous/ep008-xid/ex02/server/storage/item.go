package storage

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/rs/xid"
)

// Item represents a presisted item value as a row in storage
type Item struct {
	ID string
	// OwnerId int64
	Name        string
	Description string
	Price       float64
	CreatedOn   time.Time
	UpdatedOn   time.Time
}

type ItemsStore struct {
	items map[string]Item
}

func NewItemsStore() *ItemsStore {
	s := &ItemsStore{
		items: make(map[string]Item),
	}

	return s
}

// Save() saves an item object to the store and returns it's ID
func (s *ItemsStore) Save(name string, desc string, price float64) (string, error) {
	// validate inputs

	item := Item{
		Name:        name,
		Description: desc,
		Price:       price,
		CreatedOn:   time.Now(),
		UpdatedOn:   time.Now(),
	}

	item.ID = xid.New().String()
	s.items[item.ID] = item
	return item.ID, nil
}

// GetByID() saves an item object to the store and returns it's ID
func (s *ItemsStore) GetByID(id string) (Item, error) {
	// validate inputs
	validId, err := xid.FromString(id)
	if err != nil {
		return Item{}, fmt.Errorf("invalid item id: %v", id)
	}

	item := s.items[validId.String()]
	slog.Info("retrieved item", "id", id, "item.name", item.Name)

	return item, nil
}

