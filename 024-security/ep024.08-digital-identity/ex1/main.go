package main

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	wg sync.WaitGroup
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	host1(ch1, ch2)
	host2(ch2, ch1)

	// wait for gorountines
	wg.Wait()
}

func host1(out chan<- string, in <-chan string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		out <- "Hello, send me SECRET"
		msg := <-in
		fmt.Println("Host2 responded with:", msg)
	}()

}
func host2(out chan<- string, in <-chan string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		req := <-in
		logrus.Infof("Host2 got reqest: %v", req)
		out <- "Here is the secret"
	}()

}
