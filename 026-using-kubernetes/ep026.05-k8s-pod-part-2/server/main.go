package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/striversity/gotr/026-using-kubernetes/ep026_01/model"
	"github.com/striversity/gotr/026-using-kubernetes/ep026_01/server/store"
)

/********************************
Simple Web Server app which exposes a counter endpoint.
Methods:
| counter | POST | sets a counter value |
| counter | GET | gets the current counter value |

The server persists the counter value in a Redis database.
*/

var listenAddr = ":8080"
var redisURL = "localhost:6379"

type CounterService struct {
	db     *store.RedisStore
	router *httprouter.Router
}

func main() {
	t := os.Getenv("LISTEN_ADDRESS")
	if t != "" {
		listenAddr = t
	}

	t = os.Getenv("REDIS_URL")
	if t != "" {
		redisURL = t
	}

	cs := NewCounterService()

	err := http.ListenAndServe(listenAddr, cs.router)
	if err != nil {
		logrus.Fatalf("unable to start server: %v", err)
	}
}

func NewCounterService() *CounterService {
	cs := new(CounterService)
	cs.db = store.NewRedisStore(redisURL)

	cs.router = httprouter.New()

	cs.router.POST("/counter", cs.postCounter)
	cs.router.GET("/counter", cs.getCounter)
	return cs
}

func (cs *CounterService) postCounter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logrus.Infof("/counter POST called")
	c := new(model.Count)
	body := r.Body

	err := json.NewDecoder(body).Decode(c)
	if err != nil {
		logrus.Warnf("unable to decode counter: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "unable to decode counter: %v", err)
		return
	}

	buf, _ := json.Marshal(c)
	err = cs.db.Put("count", buf)
	if err != nil {
		logrus.Warnf("unable to save value: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error, try again later")
		return
	}
}

func (cs *CounterService) getCounter(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logrus.Infof("/counter GET called")

	buf, err := cs.db.Get("count")
	if err != nil {
		logrus.Warnf("unable to retrieve value from database: %v", err)
		fmt.Fprintf(w, "internal server error, try again later")
		return
	}

	w.Write(buf)
}
