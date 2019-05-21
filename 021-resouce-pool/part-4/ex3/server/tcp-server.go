package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync/atomic"
	"thread-pool/rp"
	"time"

	log "github.com/mgutz/logxi/v1"
)

type (
	// TCPServer listen for client on a port and accepts messages
	TCPServer struct {
		numReqs uint64
		port    string
		log     log.Logger
		s       *http.Server
	}
)

var (
	s  = rand.NewSource(time.Now().Unix())
	rn = rand.New(s)
)

// newTCPServer creates a new server to accept
func newTCPServer(port string) Server {
	srv := &TCPServer{port: port}
	srv.log = log.New("server")
	srv.log.SetLevel(log.LevelInfo)

	// TODO - create and configure http.Server
	s := &http.Server{}
	// configure http server
	s.WriteTimeout = 500 * time.Millisecond
	s.ReadTimeout = 1000 * time.Millisecond
	s.Addr = fmt.Sprintf(":%v", srv.port) // listen on all interfaces
	s.Handler = srv
	srv.s = s // store a reference to the http.Server
	return srv
}

// Start listening on srv.port and all interfaces
func (srv *TCPServer) Start() error {
	if nil == srv {
		return fmt.Errorf("Start() called on nil TCPServer object")
	}

	srv.log.Info("Starting HTTP server")
	// TODO - actually start the http server
	var err error
	go func() {
		err = srv.s.ListenAndServe()
	}()
	time.Sleep(200 * time.Millisecond)
	return err
}

// Stop listening and close all client connections
func (srv *TCPServer) Stop() {
	if nil == srv {
		return
	}

	srv.log.Info("Stopping HTTP server")
	// TODO - stop http server
	srv.s.Close()

	srv.log.Info("Messages processed:", srv.numReqs)
	rp.Stats()
}

func (srv *TCPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(&srv.numReqs, 1)

	go func(){
		srv.log.Debug("TCPServer - message from", r.RemoteAddr)
		dec := json.NewDecoder(r.Body)
		defer r.Body.Close()
		msg := rp.Alloc()
		dec.Decode(msg)
		// INFO - pretent we do some work on with the msg
		time.Sleep(10 * time.Millisecond)
		rp.Release(msg)
	}()
	
}
