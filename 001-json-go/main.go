package main

// ref: https://golang.org/pkg/encoding/json/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"./types"
)

const jsonFile = "./data/user.db.json"

func main() {
	createJsonFile()

	f, err := os.Open(jsonFile)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)

	db := types.UserDb{}
	dec.Decode(&db)
	fmt.Println(db)
}

func createJsonFile() {
	users := []types.User{
		{Id: 1, Username: "Jane Doe", Password: "please change me", Email: "janedoe@email.com"},
		{Id: 2, Username: "John Doe", Password: "change me", Email: "johndoe@email.com"},
	}

	db := types.UserDb{Users: users, Type: "Simple"}

	//	fmt.Println(users)
	var buf = new(bytes.Buffer)

	enc := json.NewEncoder(buf)
	enc.Encode(db)
	f, err := os.Create(jsonFile)
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	io.Copy(f, buf)
}
