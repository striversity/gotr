// using mgo package - https://github.com/globalsign/mgo
package main

import (
	"fmt"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	url = "localhost"
)

func main() {
	// list avaialable databases
	// ----
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Infof("Successfully connected to mongodb server at %v", url)

	dbNames, err := session.DatabaseNames()
	if err != nil {
		log.Warn(err)
	}
	for i, v := range dbNames {
		fmt.Printf("[%3v] - %v\n", i+1, v)
	}
}
