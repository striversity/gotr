package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func host2(out chan<- []byte, in <-chan []byte) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		msg := <-in
		var client HostID
		err := gobDecode(&client, msg)
		if err != nil {
			logrus.Errorf("Unable to decode client message: %v", err)
			out <- []byte("ERROR: try again")
			return
		}

		if !isTrustedHost(client) {
			m := "ALARM: Untrusted host"
			logrus.Errorf(m)
			out <- []byte(m)
			return
		}

		reply := fmt.Sprintf("Hi %v, here is the SECRET", client.Name)
		out <- []byte(reply)
	}()

}

func isTrustedHost(host HostID) bool {

	return false
}
