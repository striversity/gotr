// +build js,wasm
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()
	hostname, _ := os.Hostname()
	fmt.Printf("Hello world from host '%v' at %v\n", hostname, now)
}
