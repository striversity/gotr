package main

import (
	"flag"
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
	flag.StringVar(&address, "a", address, "gRPC server address host:port")
	flag.Parse()

	var opts []grpc.ServerOption
	// configure tls
	server := grpc.NewServer(opts...)

	model.RegisterMyMathServiceServer(server, &myMathService{})
	model.RegisterDataServiceServer(server, &myDataService{})

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to start gRPC server on address %v: %v", address, err))
	}
	server.Serve(lis)
}
