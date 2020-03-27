package main

import (
	"fmt"
	"time"
)

func main() {
	srv := newTCPServer("8080")
	srv.Start()
	d := 10 * time.Second
	fmt.Printf("Sleeping for %v\n", d)
	time.Sleep(d)
	srv.Stop()
}
