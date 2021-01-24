package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	logDir  = "/data"
	logFile = logDir + "/server.log"
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

	// add log entry
	addLogEntry(now, hostname, r.RemoteAddr)
}

func addLogEntry(t time.Time, lHost string, rHost string) {
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0644))
	if err != nil {
		logrus.Fatal(err)
	}

	defer f.Close()

	fmt.Fprintf(f, "%s\tAPI Server %v: Connection from - %v\n", t.Format(time.RFC3339), lHost, rHost)
}
