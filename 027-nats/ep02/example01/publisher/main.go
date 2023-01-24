package main

import (
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
	for {
		nc.Publish("intros", []byte("Hello World!"))
		count++
		log.Printf("sent message %v", count)
		time.Sleep(1 * time.Second)
	}
}
