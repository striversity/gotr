package main

import (
	"net/http"

	"github.com/charmbracelet/log"
)

// itemsHandler is a HTTP handler for the items.
func itemsHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	log.Infof("Received a %s request for items", method)
	switch method {
	case "GET":
		w.Write([]byte("GET"))
	case "POST":
		w.Write([]byte("POST"))
	case "PUT":
		w.Write([]byte("PUT"))
	case "DELETE":
		w.Write([]byte("DELETE"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/items", itemsHandler)

	log.Info("Starting server on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}

}
