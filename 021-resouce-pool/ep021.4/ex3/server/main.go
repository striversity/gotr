package main

import (
	"fmt"
	"time"

	log "github.com/mgutz/logxi/v1"
)

func main() {
	srv := newTCPServer("8080")
	err := srv.Start()
	if err != nil {
		log.Error("Failed to start TCPServer", err)
		return
	}
	d := 50 * time.Second
	fmt.Printf("Sleeping for %v\n", d)
	time.Sleep(d)
	srv.Stop()
}
