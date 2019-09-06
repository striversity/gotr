package main

import (
	"fmt"
	"log"
	"mms/model"
	"net"

	"google.golang.org/grpc"
)

var (
	address = "localhost:8080"
)

func main() {
	var opts []grpc.ServerOption
	// configure tls
	server := grpc.NewServer(opts...)

	model.RegisterMyMathServiceServer(server, &myMathService{})

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to start gRPC server on address %v: %v", address, err))
	}
	server.Serve(lis)
}
