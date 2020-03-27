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
	fmt.Printf("Message size: %v\n", len(msgToServer.Data))

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

	msgIn := &model.ClientReq{}
	err = proto.Unmarshal(buf, msgIn)
	if err != nil {
		log.Printf("Unable to decode Protobuf message: %v", err)
		return
	}

	fmt.Printf("Protobuf message: ID: %v, Type: %v, size: %v\n",
		msgIn.ID, msgIn.Type, len(msgIn.Data))
}

func testGob(msg *model.ClientReq) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(msg)
	fmt.Printf("GOB message size: %v\n", buf.Len())

	msgIn := &model.ClientReq{}
	dec := gob.NewDecoder(buf)
	err := dec.Decode(msgIn)
	if err != nil {
		log.Printf("Unable to decode GOB message: %v", err)
		return
	}

	fmt.Printf("GOB message: ID: %v, Type: %v, size: %v\n",
		msgIn.ID, msgIn.Type, len(msgIn.Data))
}

func testJson(msg *model.ClientReq) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.Encode(msg)
	fmt.Printf("JSON message size: %v\n", buf.Len())

	msgIn := &model.ClientReq{}
	dec := json.NewDecoder(buf)
	err := dec.Decode(msgIn)
	if err != nil {
		log.Printf("Unable to decode GOB message: %v", err)
		return
	}

	fmt.Printf("JSON message: ID: %v, Type: %v, size: %v\n",
		msgIn.ID, msgIn.Type, len(msgIn.Data))
}

func testXml(msg *model.ClientReq) {
	buf := &bytes.Buffer{}
	enc := xml.NewEncoder(buf)
	enc.Encode(msg)
	fmt.Printf("XML message size: %v\n", buf.Len())

	msgIn := &model.ClientReq{}
	dec := xml.NewDecoder(buf)
	err := dec.Decode(msgIn)
	if err != nil {
		log.Printf("Unable to decode GOB message: %v", err)
		return
	}

	fmt.Printf("XML message: ID: %v, Type: %v, size: %v\n",
		msgIn.ID, msgIn.Type, len(msgIn.Data))
}
