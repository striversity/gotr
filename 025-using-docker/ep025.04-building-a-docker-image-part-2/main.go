package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	port = "8080"
)

func main() {
	flag.StringVar(&port, "p", port, "this is the port number")
	flag.Parse()
	logrus.Infof("Starting HTTP server")
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		logrus.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello %v, my name is '%v' at %v\n", r.RemoteAddr, hostname, now)
}
