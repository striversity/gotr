package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/striversity/gotr/026-using-kubernetes/ep026_01/model"
)

var apiURL = "http://localhost:8080/counter"
var rg = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	t := os.Getenv("API_URL")
	if t != "" {
		apiURL = t
	}

	for {
		postCount()
		d := rg.Intn(5000)
		time.Sleep(time.Millisecond * time.Duration(d))
	}
}

func postCount() {
	c := model.Count{
		Counter: rg.Intn(1e4),
	}

	b, err := json.Marshal(c)
	if err != nil {
		logrus.Warnf("unable to encode count: %v", err)
		return
	}

	buf := bytes.NewBuffer(b)
	_, err = http.Post(apiURL, "application/json", buf)
	if err != nil {
		logrus.Errorf("unable to post count: %v", err)
	}

	logrus.Infof("posted count value %v\n", c.Counter)
}
