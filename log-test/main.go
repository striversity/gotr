package main

import log "github.com/sirupsen/logrus"

func main() {
	log.Info("This is a log line")
	log.Warn("Another log line")
	log.Fatal("This is really bad, exiting...")
}
