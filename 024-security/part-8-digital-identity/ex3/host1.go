package main

import "fmt"

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
