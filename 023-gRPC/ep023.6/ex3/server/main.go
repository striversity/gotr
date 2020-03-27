package main

import (
	"flag"
	"fmt"
	"log"
	"mms/model"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	address = "vee.mv.lorrev.org:8080"
)

func main() {
	flag.StringVar(&address, "a", address, "gRPC server address host:port")
	flag.Parse()

	var opts []grpc.ServerOption
	// configure tls
	creds, _ := credentials.NewServerTLSFromFile("../cert.pem", "../key.pem")

	opts = append(opts, grpc.Creds(creds))

	server := grpc.NewServer(opts...)

	model.RegisterMyMathServiceServer(server, &myMathService{})
	model.RegisterDataServiceServer(server, &myDataService{})

	log.Printf("gRPC Server listening on %v\n", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to start gRPC server on address %v: %v", address, err))
	}
	server.Serve(lis)
}
