// using mgo package - https://github.com/globalsign/mgo
package main

import (
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	url = "localhost"
)

func main() {
	// connecting to mongodb server
	// ----
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Infof("Successfully connected to mongodb server at %v", url)
}
