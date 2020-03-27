package main

import (
	"benc/model"
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	msgToServer := &model.ClientReq{ID: 1, ReqType: 3, Size: 255}
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)

	enc.Encode(msgToServer)

	fmt.Printf("buf size: %v\n", buf.Len())
}
