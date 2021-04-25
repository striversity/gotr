package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

var (
	rg = rand.New(rand.NewSource(time.Now().Unix()))
)

func main() {
	url := "nats://localhost:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		logrus.Fatal(err)
	}

	defer nc.Close()

	for i := 0; i < 1e5; i++ {
		s := fmt.Sprintf("Message %v: data: %v", i, rg.Intn(10000))

		nc.Publish("events.old", []byte(s))
	}

}
