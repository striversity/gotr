package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	address = "vee.mv.lorrev.org:8080"
)

func main() {
	flag.StringVar(&address, "a", address, "gRPC server address host:port")
	flag.Parse()

	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServeTLS(address, "../cert.pem", "../key.pem", nil)
	if err != nil {
		log.Fatalf("Unable to start Server on %v: %v\n", address, err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got connection from client: %v\n", r.RemoteAddr)

	w.Header().Set("content-type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello there, nice to meet you %v\n", r.RemoteAddr)
}
