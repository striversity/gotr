package main

import (
	"awesome/model"
	"encoding/json"
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

	sub, err := nc.QueueSubscribe("intros", "zip1", processMsg)
	if err != nil {
		log.Fatalf("can't subscribe to NATS queue 'zip1': %v", err)
	}
	defer sub.Unsubscribe()

	time.Sleep(1 * time.Hour)
}

func processMsg(m *nats.Msg) {
	pl := &model.Payload{}
	json.Unmarshal(m.Data, pl)
	replyData := fmt.Sprintf("ack message # %v", pl.Count)
	m.Respond([]byte(replyData))
	fmt.Printf("I got a message: %s, count %v\n", pl.Data, pl.Count)
}
