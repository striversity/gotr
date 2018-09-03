/*
Simple example to demonstrate GO's SQL driver talking to PostgreSQL Server.
*/
package main

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func main() {
	log.Info("Connecting to SQL Db...")

	connStr := "user=another dbname=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var id int
	var name, username, pswd = "Peter Jones", "pj@email.com", "password"
	qr, err := db.Exec(`INSERT INTO users (name, username, password) 
	   VALUES ($1, $2, $3);`, name, username, pswd)
	if err != nil {
		log.Warn(err)
	}
	fmt.Printf("Query Result: %v\n", qr)

	rows, err := db.Query("SELECT id, name, username FROM users")
	for rows.Next() {
		rows.Scan(&id, &name, &username)
		fmt.Printf("Got: Id: %v, Name: %v, Username: %v\n", id, name, username)
	}
}
