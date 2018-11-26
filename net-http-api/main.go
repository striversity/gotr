package main

import (
	"log"
	"net/http"
	"github.com/striversity/go-on-the-run/net-http-api/api/user"
)

func main() {
	// register RESTful endpoint handler for '/users/'
	http.Handle("/users/", &user.UserAPI{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
