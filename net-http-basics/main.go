package main

import (
	"fmt"
	"log"
	"net/http"
)

const usersApiResp = `
<html>
<body>
<p>Hi, thanks for calling my /users API with HTTP Method '%v'
<p>This is the %v call to this API
</body>
</html>
`

var userCounter int

type reportCounter struct {
	counter int
}

func main() {
	http.HandleFunc("/users", usersHandleFunc)
	var rc reportCounter
	http.Handle("/reports/", &rc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (rc *reportCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We got a request on /reports")
	rc.counter++
	s := fmt.Sprintf("/reports API call count: %v", rc.counter)
	fmt.Fprint(w, s)
}

func usersHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We got a request on /users")
	userCounter++
	s := fmt.Sprintf(usersApiResp, r.Method, userCounter)
	fmt.Fprint(w, s)
}
