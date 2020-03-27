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
	// list documents in selected database collections
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
	fmt.Printf("Collections in db '%v':\n", dbName)
	cols, err := db.CollectionNames()
	if err != nil {
		return
	}

	for _, c := range cols {
		fmt.Printf("[%v]\n", c)
		listDocs(db, c)
	}
}

func listDocs(db *mgo.Database, col string) {
	coll := db.C(col)
	if coll == nil {
		return
	}

	var result []map[string]interface{}
	coll.Find(nil).All(&result)
	for i, d := range result {
		fmt.Printf("\tDoc%3v - %v\n", i+1, d)
	}
}
