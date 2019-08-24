// nested goroutine using context.WithValue()
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

const (
	numWorkers = 3
)

type (
	topSecretKey string
)

func main() {
	parent := context.Background()

	k := topSecretKey("go now")
	v := "time to go!"
	valueContext := context.WithValue(parent, k, v)

	t := 1 * time.Second
	ctx, cancel := context.WithTimeout(valueContext, t)
	defer cancel()

	// start a number of goroutines and wait for them to complete
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, fmt.Sprintf("%v", i))
	}

	<-ctx.Done()
}

// worker does some work for a random duration of time
func worker(ctx context.Context, id string) {
	k := topSecretKey("go now")
	v := ctx.Value(k)

	d := time.Duration(r.Int31n(5000)) * time.Millisecond
	fmt.Printf("Worker %v started for %v with valule: %v\n", id, d, v)

	t := time.Now().Add(500 * time.Millisecond)
	ctx2, cancel := context.WithDeadline(ctx, t)
	go subWorker(ctx2, id+".1")
	defer cancel()

	select {
	case <-time.After(d):
	case <-ctx.Done():
		return
	}

	fmt.Printf("Worker %v completed after %v\n", id, d)
}

// subWorker does some work for a random duration of time
func subWorker(ctx context.Context, id string) {
	d := time.Duration(r.Int31n(5000)) * time.Millisecond
	fmt.Printf("subWorker %v started for %v\n", id, d)

	select {
	case <-time.After(d):
	case <-ctx.Done():
		return
	}

	fmt.Printf("subWorker %v completed after %v\n", id, d)
}
