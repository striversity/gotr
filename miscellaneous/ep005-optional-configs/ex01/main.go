package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	"github.com/striversity/gotr/misc-ep005/ex01/streamer"
)

func main() {
	db, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
	if err != nil {
		logrus.Fatal(err)
	}
	query := `SELECT * FROM products;`
	interval := 30 * time.Minute
	dbStreamer := streamer.DBStreamConfig{
		db: db, query: query, interval: interval,
	}

	dbStreamer.Field = "hello" // no desirable
	dbStreamer.Start()
	// ...
	dbStreamer.Stop()
}
