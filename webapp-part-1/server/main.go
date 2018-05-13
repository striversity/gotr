package main

import (
	"log"
	"net/http"

	"esurelabs.homelinux.com/verrol/go-on-the-run/webapp-part-1/server/api/user"
)

func main() {
	// register RESTful endpoint handler for '/users/'
	http.Handle("/users/", &user.UserAPI{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
