package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("tls://127.0.0.1:4222",
		nats.UserInfo("auser2", "auser2"))
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	go func() {
		for {
			now := time.Now()
			msg := fmt.Sprintf("Hello, the time is %v",
				now.Format("15:04:05"))
			nc.Publish("foo", []byte(msg))
			time.Sleep(1 * time.Second)
		}
	}()

	// keep the program running forever
	select {}
}
