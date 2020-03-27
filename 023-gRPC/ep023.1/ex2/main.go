// goroutine cancellation using channel
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	s  = rand.NewSource(time.Now().Unix())
	r  = rand.New(s)
	wg = &sync.WaitGroup{}
)

const (
	numWorkers = 3
)

func main() {
	ch := make(chan bool, numWorkers)

	// start a number of goroutines and wait for them to complete
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ch, i)
	}

	<-time.After(1 * time.Second)
	for i := 0; i < numWorkers; i++ {
		ch <- true
	}
	wg.Wait()
}

// worker does some work for a random duration of time
func worker(ch chan bool, id int) {
	defer wg.Done()

	d := time.Duration(r.Int31n(5000)) * time.Millisecond
	fmt.Printf("Worker %v started for %v\n", id, d)

	select {
	case <-time.After(d):
	case <-ch:
		return
	}

	fmt.Printf("Worker %v completed after %v\n", id, d)
}
