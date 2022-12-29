package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var hostname string

func main() {
	hostname, _ = os.Hostname()

	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Fatalf("unable to listen: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("API server %v at %v", hostname, time.Now())
	logrus.Info(msg)
	fmt.Fprintf(w, "%s\n", msg)
}
