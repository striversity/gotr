package main

import (
	"benc/model"
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	// client side
	msgToServer := &model.ClientReq{ID: 1, ReqType: 3, Size: 255}
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(msgToServer)

	// TODO send encoded message to server
	// ...
	fmt.Printf("buf size: %v\n", buf.Len())

	// server side
	msgFromClient := &model.ClientReq{}
	gob.NewDecoder(buf).Decode(msgFromClient)
	fmt.Printf("msgFromClient: %v\n", msgFromClient)
}
