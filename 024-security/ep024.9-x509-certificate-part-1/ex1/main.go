package main

import (
	"sec/ex9/issuer"
	"sync"
)

type (
	Request struct {
		From    *issuer.SignedHostID
		Message string
	}
)

var (
	wg sync.WaitGroup
)

func main() {
	ch1 := make(chan []byte)
	ch2 := make(chan []byte)

	host1(ch1, ch2)
	host2(ch2, ch1, issuer.PublicKey())

	// wait for gorountines
	wg.Wait()
}
