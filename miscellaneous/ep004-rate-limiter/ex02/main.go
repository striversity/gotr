package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/time/rate"
)

var (
	rg = rand.New(rand.NewSource(time.Now().Unix()))
)

func main() {
	r := rate.Every(2 * time.Second)
	lim := rate.NewLimiter(r, 3)

	for i := 0; i < 10; i++ {
		n := lim.Burst()
		lim.WaitN(context.Background(), n)
		for j := 0; j < n; j++ {
			callExternal()
		}
	}
}

func callExternal() {
	fmt.Printf("request to external at %v\n", time.Now())
	d := time.Duration(rg.Int31n(500)) * time.Millisecond
	time.Sleep(d)
}
