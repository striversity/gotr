package main

import (
	"encoding/csv"
	"encoding/json"
	"os"

	"./types"
	log "github.com/Sirupsen/logrus"
)

const (
	jsonFile = "./data/user.db.json"
	csvFile  = "./data/user.db.csv"
)

func main() {
	db, err := readJSONFile(jsonFile)
	if nil != err {
		log.Fatalln(err)
	}

	f, err := os.Create(csvFile)
	if nil != err {
		log.Fatalln(err)
	}

	defer f.Close()

	w := csv.NewWriter(f)

	w.Write(types.GetHeader())
	for _, user := range db.Users {
		ss := user.EncodeAsStrings()
		w.Write(ss)
	}

	w.Flush()

	err = w.Error()
	if nil != err {
		log.Fatalln(err)
	}
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
