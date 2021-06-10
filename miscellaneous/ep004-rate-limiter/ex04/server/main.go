package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

var (
	rg  = rand.New(rand.NewSource(time.Now().Unix()))
	r   = rate.Every(2 * time.Second)
	lim = rate.NewLimiter(r, 3)
)

func main() {

	logrus.Info("Starting HTTP Server on port :8080")

	http.HandleFunc("/api", apiHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Error(err)
	}
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if lim.Allow() {
		d := time.Duration(rg.Int31n(500)) * time.Millisecond
		time.Sleep(d)
		fmt.Fprintf(w, "The time is %v\n", time.Now())
		return
	}

	logrus.Warnf("too many calls, not handling")
	w.WriteHeader(http.StatusBadRequest)
}
