package main

import (
	"fmt"
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

	nc.Subscribe("intros", func(m *nats.Msg) {
		fmt.Printf("I got a message: %s\n", string(m.Data))
	})

	time.Sleep(1 * time.Hour)
}
