// waiting for all goroutines to complete after a random time between 0-5000 milliseconds
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
	// start a number of goroutines and wait for them to complete
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()
}

// worker does some work for a random duration of time
func worker(id int) {
	d := time.Duration(r.Int31n(5000)) * time.Millisecond
	fmt.Printf("Worker %v started for %v\n", id, d)
	time.Sleep(d)
	fmt.Printf("Worker %v completed after %v\n", id, d)
	wg.Done()
}
