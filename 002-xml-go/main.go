package main

// ref: https://golang.org/pkg/encoding/xml/

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"

	"./types"
)

const (
	jsonFile = "./data/user.db.json"
	xmlFile  = "./data/user.db.xml"
)

func main() {
	db, err := readJSONFile(jsonFile)
	if nil != err {
		log.Fatalln(err)
	}

	f, err := os.Create(xmlFile)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()

	xmlEnc := xml.NewEncoder(f)
	xmlEnc.Encode(db)
}

func readJSONFile(s string) (db *types.UserDb, err error) {
	f, err := os.Open(s)
	if nil != err {
		return nil, err
	}
	defer f.Close()

	var dec = json.NewDecoder(f)

	db = new(types.UserDb)
	dec.Decode(db)

	return
}
