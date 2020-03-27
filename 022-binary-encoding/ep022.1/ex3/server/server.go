package main

type (
	// Server can be started and stopped
	Server interface {
		Start() error
		Stop()
	}
)
