package main

import (
	"awesome/model"
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}
	defer nc.Close()

	count := 0
	pl := &model.Payload{
		Data: "Hello, world!",
	}

	for {
		pl.Count = count
		data, _ := json.Marshal(pl)
		reply, err := nc.Request("intros", data, 500*time.Millisecond)
		time.Sleep(1 * time.Second)
		if err != nil {
			log.Printf("error sending message count = %v, err: %v", count, err)
			continue
		}
		count++
		log.Printf("sent message %v, got reply %v", count, string(reply.Data))
	}
}
