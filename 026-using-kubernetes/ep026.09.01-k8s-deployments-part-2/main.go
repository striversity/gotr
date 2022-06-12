package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	hostname, _ := os.Hostname()
	version := os.Getenv("VERSION")

	if version == "v4.0" {
		panic("bad version")
	}
	
	for {
		time.Sleep(2 * time.Second)
		fmt.Printf("Running version %v on host %v\n", version, hostname)
	}
}
