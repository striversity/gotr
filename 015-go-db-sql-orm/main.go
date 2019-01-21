/*
Using the GORM ORM in Go.
*/
package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

const (
	connStr = "user=another dbname=postgres host=localhost sslmode=disable"
)

type (
	User struct {
		gorm.Model
		Name     string
		Username string `gorm:"not null"`
		Password string `gorm:"not null"`
		Messages []Message
	}
	Profile struct {
		TZ *time.Location
	}
	Message struct {
		Body   string `gorm:"not null`
		User   User
		UserID uint
		gorm.Model
	}
)

func main() {
	log.Info("Connecting to SQL Db...")

	db, err := gorm.Open("postgres", connStr)
	// db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Profile{})
	db.AutoMigrate(&Message{})

	var user = &User{Name: "Peter Jones", Username: "pj@email.com", Password: "password"}
	// CRUD = create, retrieve, update, and delete
	if err := db.Model(&User{}).Create(user).Error; err != nil {
		log.Warn(err)
	}
	users, _ := db.Model(&User{}).Find(user).Rows()
	for users.Next() {
		u := new(User)
		db.ScanRows(users, u)
		fmt.Printf("Got: Name: %v, Username: %v\n", u.Name, u.Username)
	}

	mesgs := []*Message{
		&Message{Body: "My message 1"},
		&Message{Body: "My message 2"},
		&Message{Body: "My message 3"},
	}

	db.Model(user).Find(user)
	for _, m := range mesgs {
		db.Model(user).Association("Messages").Append(m)
	}
}
