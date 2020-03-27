package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	server = "localhost:8080"
)

func main() {
	flag.StringVar(&server, "s", server, "gRPC server server host:port")
	flag.Parse()

	url := fmt.Sprintf("http://%v/hello", server)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Unable to connect to %v\n", server)
	}

	defer resp.Body.Close()

	fmt.Printf("Response from %s:\n", server)
	io.Copy(os.Stdout, resp.Body)
	fmt.Println()
}
