package main

import "flag"

var (
	defaultServer = ":8080"
)

func main() {
	var url string
	flag.StringVar(&url, "s", defaultServer, "server host:port")
	flag.Parse()

	submitRequests(url)
}
