package main

import (
	"log"
	"log/slog"

	"github.com/google/uuid"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type Order struct {
	OrderId  string `json:"orderId"`
	Username string `json:"username"`
	CartId   string `json:"cartId"`
}

var (
	orderQueue = make(chan Order, 10)
)

func main() {
	app := pocketbase.New()

	// run this hook for updates to the 'carts' collection
	app.OnRecordAfterUpdateRequest("carts").Add(processCartOrder)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func processCartOrder(e *core.RecordUpdateEvent) error {
	rec := e.Record
	slog.Info("record updated", "collection", "carts", "recordId", rec.Id)

	if !rec.GetBool("ordered") {
		return nil
	}

	// send record for downstream to another service for fulfillment/processing
	newOrder := Order{
		OrderId:  uuid.NewString(),
		Username: rec.GetString("user"),
		CartId:   rec.Id,
	}

	orderQueue <- newOrder
	slog.Info("sent order to fulfillment service", "order", newOrder)

	return nil
}
