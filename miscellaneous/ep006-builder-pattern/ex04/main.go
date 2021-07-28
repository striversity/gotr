package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"

	"github.com/striversity/gotr/misc-ep006/ex04/streamer"
)

func main() {
	db, err := sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
	if err != nil {
		logrus.Fatal(err)
	}
	query := `SELECT * FROM products;`

	dbStreamer, err := streamer.NewBuilder().WithQuery(query).WithInterval(0 * time.Second).WithDb(nil).Build()
	if err != nil {
		logrus.Fatal(err)
	}

	_ = db
	dbStreamer.Start()
	time.Sleep(20 * time.Second)
	dbStreamer.Stop()
}
