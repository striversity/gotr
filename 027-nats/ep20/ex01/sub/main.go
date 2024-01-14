package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("tls://127.0.0.1:4222",
		nats.UserInfo("auser1", "auser1"))
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	// Subscribe to the j"foo" subject and print all messages
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Println(string(m.Data))
	})

	// Keep the application running
	select {}
}
