package main

import (
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	rg = rand.New(rand.NewSource(time.Now().Unix()))
)

func main() {
	client := http.DefaultClient

	for i := 0; i < 10; {
		resp, err := client.Get("http://localhost:8080/api")
		if err != nil || resp.StatusCode != http.StatusOK{
			logrus.Errorf("request failed: %v", err)
			d := time.Duration(rg.Int31n(500)) * time.Millisecond
			time.Sleep(d)
		} else {
			i++
			io.Copy(os.Stdout, resp.Body)
			resp.Body.Close()
		}
	}
}
