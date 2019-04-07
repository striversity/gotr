package main

import "flag"

var (
	defaultServer = "http://localhost:8080"
)

func main() {
	var url string
	flag.StringVar(&url, "s", defaultServer, "server in the form http://host:port")
	flag.Parse()

	submitRequests(url)
}
