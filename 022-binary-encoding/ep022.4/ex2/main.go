package main

import (
	"benc/model"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	msgToServer := &model.ClientReq{ID: 1, Type: model.ReqType_ADD}
	msgToServer.Data = []byte("5 + 7")

	testXml(msgToServer)
	testJson(msgToServer)
	testGob(msgToServer)
	testProtobuf(msgToServer)
}

func testProtobuf(msg *model.ClientReq) {
	buf, err := proto.Marshal(msg)
	if err != nil {
		log.Printf("Unable to encode Protobuf message: %v", err)
		return
	}

	fmt.Printf("Protobuf message size: %v\n", len(buf))
}

func testGob(msg *model.ClientReq) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(msg)
	fmt.Printf("GOB message size: %v\n", buf.Len())
}

func testJson(msg *model.ClientReq) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.Encode(msg)
	fmt.Printf("JSON message size: %v\n", buf.Len())
}

func testXml(msg *model.ClientReq) {
	buf := &bytes.Buffer{}
	enc := xml.NewEncoder(buf)
	enc.Encode(msg)
	fmt.Printf("XML message size: %v\n", buf.Len())
}
