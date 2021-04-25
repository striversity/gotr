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

	sub, _ := nc.Subscribe("events.*", func(msg *nats.Msg) {
		fmt.Printf("message recieved on subject: %v, data: %v\n", msg.Subject,
			string(msg.Data))
	})

	// sub, _ := nc.SubscribeSync("events.*")
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	// for {
	// 	if msg, _ := sub.NextMsg(10 * time.Second); msg != nil {
	// 		fmt.Printf("message recieved on subject: %v, data: %v\n",
	// 			msg.Subject,
	// 			string(msg.Data))
	// 	} else {
	// 		break
	// 	}
	// }

	time.Sleep(1 * time.Minute)

	sub.Unsubscribe()

}
