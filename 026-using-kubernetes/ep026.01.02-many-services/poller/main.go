package main

import (
	"encoding/json"
	"fmt"
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
		c, err := getCount()
		if err == nil {
			fmt.Printf("Count from server is %v\n", c)
		}

		d := rg.Intn(5000)
		time.Sleep(time.Millisecond * time.Duration(d))
	}
}

func getCount() (int, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		logrus.Warnf("unable to get count: %v", err)
		return 0, err
	}

	c := new(model.Count)
	err = json.NewDecoder(resp.Body).Decode(c)
	if err != nil {
		logrus.Errorf("unable to decode count: %v", err)
		return 0, err
	}

	return c.Counter, nil

}
