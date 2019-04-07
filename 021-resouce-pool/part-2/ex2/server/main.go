package main

func main() {
	srv := newTCPServer("8080")
	srv.Start()
}
