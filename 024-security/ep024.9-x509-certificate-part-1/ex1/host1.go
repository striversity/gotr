package main

import (
	"fmt"
	"sec/ex9/issuer"
	"sec/ex9/serialize"

	"github.com/sirupsen/logrus"
)

func host1(out chan<- []byte, in <-chan []byte) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		
		myID, err := issuer.NewSignedHostID("host1", []string{"10.10.1.2"})
		if err != nil {
			logrus.Error(err)
		}

		req := Request{From: myID, Message: "Send me message"}

		out <- serialize.GobEncode(req)
		msg := <-in
		fmt.Println("Host2 responded with:", string(msg))
	}()

}
