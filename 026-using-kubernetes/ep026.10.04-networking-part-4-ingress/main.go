package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var hostname string
var appVersion = os.Getenv("VERSION")

func main() {
	hostname, _ = os.Hostname()
	if appVersion == "" {
		appVersion = "v0.0.1"
	}

	http.HandleFunc("/srv2/posts", postsHandler)
	http.HandleFunc("/comments", commentsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Fatalf("unable to listen: %v", err)
	}
}

func commentsHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("got request for %v on API server %v, version %v at %v",
		r.RequestURI, hostname, appVersion, time.Now())
	logrus.Info(msg)
	fmt.Fprintf(w, "%s\n", msg)
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("got request for %v on API server %v, version %v at %v",
		r.RequestURI, hostname, appVersion, time.Now())
	logrus.Info(msg)
	fmt.Fprintf(w, "%s\n", msg)
}
