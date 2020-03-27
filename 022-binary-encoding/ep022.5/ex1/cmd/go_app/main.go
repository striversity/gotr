package main

import (
	"benc/model"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	cr1 := &model.SearchReq{}

	fmt.Printf("cr1: %v\n", cr1)

	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Unable to read file: %v", os.Args[1])
	}

	proto.Unmarshal(buf, cr1)

	fmt.Printf("Req.q: %v\n", cr1.GetQ())
	for k, v := range cr1.GetParams() {
		fmt.Printf("Req key: %v -> value: %v\n", k, v)
	}
}
