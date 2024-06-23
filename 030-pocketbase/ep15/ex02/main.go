package main

import (
	"log"
	"log/slog"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
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
	slog.Info("record updated", "collection", "carts", "recordId", e.Record.Id)
	return nil
}
