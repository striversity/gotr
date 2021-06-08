package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main() {
	url := "nats://localhost:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		logrus.Fatal(err)
	}
	defer nc.Close()

	count := 0
	sub, _ := nc.Subscribe("events.*", func(msg *nats.Msg) {
		count++
		fmt.Printf("message recieved on subject: %v, data: %v\n", msg.Subject,
			string(msg.Data))
		if msg.Reply != "" {
			msg.Respond([]byte("got it"))
		}
	})

	defer sub.Unsubscribe()

	for {
		old := count
		time.Sleep(5 * time.Second)
		if old == count {
			break
		}
	}

	fmt.Printf("processed %v messages\n", count)
}
