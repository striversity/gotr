package main

import (
	"fmt"

	log "github.com/mgutz/logxi/v1"
)

type (
	// TCPServer listen for client on a port and accepts messages
	TCPServer struct {
		numReqs uint64
		port    string
		log     log.Logger
	}
)

// newTCPServer creates a new server to accept
func newTCPServer(port string) Server {
	srv := &TCPServer{port: port}
	srv.log = log.New("server")
	srv.log.SetLevel(log.LevelDebug)

	// TODO - create and configure http.Server
	return srv
}

// Start listening on srv.port and all interfaces
func (srv *TCPServer) Start() error {
	if nil == srv {
		return fmt.Errorf("Start() called on nil TCPServer object")
	}

	srv.log.Info("Starting HTTP server")
	// TODO - actually start the http server
	return nil 
}

// Stop listening and close all client connections
func (srv *TCPServer) Stop() {
	if nil == srv {
		return
	}

	srv.log.Info("Stopping HTTP server")
	// TODO - stop http server
}
