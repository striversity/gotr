package main

import (
	"benc/model"
	"fmt"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	msgToServer := &model.ClientReq{ID: 1, Type: model.ReqType_ADD}
	msgToServer.Data = []byte("5 + 7")

	buf, err := proto.Marshal(msgToServer)
	if err != nil {
		log.Printf("Unable to encode Protobuf message: %v", err)
		return
	}

	fmt.Printf("buf size: %v\n", len(buf))
}
