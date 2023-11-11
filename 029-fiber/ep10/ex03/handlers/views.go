package handlers

import (
	"awesome/ex03/services"

	"github.com/gofiber/fiber/v2"
)

type ItemViews struct {
	repo *services.ItemRepository
}

func NewItemView(ir *services.ItemRepository) *ItemViews {
	return &ItemViews{
		repo: ir,
	}
}

func (h *ItemViews) ListItems(ctx *fiber.Ctx) error {
	model := struct {
		Items []services.Item
	}{
		Items: h.repo.GetItems(),
	}
	
	return ctx.Render("views/items/list.html", model)
}
