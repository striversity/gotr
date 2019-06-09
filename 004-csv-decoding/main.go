package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"./types"
	log "github.com/Sirupsen/logrus"
)

const (
	csvDb = "./data/user.db.csv"
)

func main() {
	log.Info("CSV Decoding")

	f, err := os.Open(csvDb)
	if nil != err {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Read() // ignore header, since we already know it

	for {
		csvRecord, err := r.Read()
		if nil == err {
			process(csvRecord)
		} else if io.EOF == err {
			break
		} else {
			log.Fatal(err)
		}
	}
}

func process(ss []string) {
	u := &types.User{}
	u.FromCSV(ss)
	fmt.Println(u.FirstName, u.Email)
}
