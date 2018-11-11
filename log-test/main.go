package main

import (
	log2 "github.com/mgutz/logxi/v1"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("This is a log line")
	log.Warn("Another log line")
	log.Error("This is really bad")
	log2.Info("This is a log line")
	log2.Warn("Another log line")
	// log2.Error("This is really bad")
}
