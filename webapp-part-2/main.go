package main

import (
	"log"
	"net/http"

	"github.com/striversity/go-on-the-run/webapp-part-2/server/api/user"
)

func main() {
	// register static files handle '/index.html -> client/index.html'
	http.Handle("/", http.FileServer(http.Dir("client")))
	// register RESTful endpoint handler for '/users/'
	http.Handle("/users/", &user.UserAPI{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
