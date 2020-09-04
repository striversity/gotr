package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

type (
	// HostID contains information that to uniquely identifies a host
	HostID struct {
		Name      string
		IPAddress []string
	}
)

var (
	wg sync.WaitGroup
)

func main() {
	ch1 := make(chan []byte)
	ch2 := make(chan []byte)

	host1(ch1, ch2)
	host2(ch2, ch1)

	// wait for gorountines
	wg.Wait()
}

func host1(out chan<- []byte, in <-chan []byte) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		myID := HostID{Name: "host1", IPAddress: []string{"10.10.1.2"}}

		out <- gobEncode(myID)
		msg := <-in
		fmt.Println("Host2 responded with:", string(msg))
	}()

}
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

		reply := fmt.Sprintf("Hi %v, here is the SECRET", client.Name)
		out <- []byte(reply)
	}()

}

func gobDecode(v interface{}, buf []byte) error {
	if v == nil || buf == nil {
		return fmt.Errorf("Invalid paraemters, v nor buf cannot be nil")
	}

	b := bytes.NewReader(buf)
	dec := gob.NewDecoder(b)

	return dec.Decode(v)
}

func gobEncode(v interface{}) []byte {
	b := &bytes.Buffer{}
	enc := gob.NewEncoder(b)
	enc.Encode(v)
	return b.Bytes()
}
