package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	username := "a"
	password := "a.very.secure.password!"
	hostport := "localhost:4222"

	url := fmt.Sprintf("nats://%s:%s@%s", username, password, hostport)
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for i := 0; i < 10; i++ {
		msg, err := nc.Request("time.of.day", []byte(""), 1*time.Second)
		if err == nil {
			fmt.Printf("INFO - Got reply - %s\n", string(msg.Data))
		}
		
		time.Sleep(1 * time.Second)
	}
}
