// using mgo package - https://github.com/globalsign/mgo
package main

import (
	"fmt"
	"os"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	url = "localhost"
)

func main() {
	dbName := "test"
	if 1 == len(os.Args) {
		log.Warnf("No db specified, using '%v'", dbName)
	} else {
		dbName = os.Args[1]
	}
	// list collections in selected database
	// ----
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Infof("Successfully connected to mongodb server at %v", url)

	db := session.DB(dbName)
	if db == nil {
		log.Errorf("db '%v' not found, exiting...", dbName)
		return
	}

	// iterate collections
	cols, err := db.CollectionNames()
	if err != nil {
		log.Warnf("No collections in db '%v'", dbName)
	}

	fmt.Printf("Collections in db '%v':\n", dbName)
	for i, v := range cols {
		fmt.Printf("[%3v] - %v\n", i+1, v)
	}
}
