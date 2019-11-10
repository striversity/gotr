package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mms/auth"
	"mms/model"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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

	uic := grpc.UnaryInterceptor(myUnaryServerInterceptor)
	opts = append(opts, uic)

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

func myUnaryServerInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for k, v := range md {
			log.Printf("%v - key: %v --> value: %v", info.FullMethod, k, v)
		}
	}

	return handler(ctx, req)
}

func authorized2(ctx context.Context, method string) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("%v - No metadata provided", method)
		return false, fmt.Errorf("%v - No metadata provided", method)
	}

	tokens := md.Get(auth.MethodKey2)
	if len(tokens) == 0 {
		return false, fmt.Errorf("%v - No auth info provided", method)
	}

	if tokens[0] == auth.MethodValue2 {
		return true, nil
	}

	return false, fmt.Errorf("%v - Not authenticated/authorized", method)
}

func authorized(ctx context.Context, method string) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("%v - No metadata provided", method)
		return false, fmt.Errorf("%v - No metadata provided", method)
	}

	tokens := md.Get(auth.MethodKey1)
	if len(tokens) == 0 {
		return false, fmt.Errorf("%v - No auth info provided", method)
	}

	if tokens[0] == auth.MethodValue1 {
		return true, nil
	}

	return false, fmt.Errorf("%v - Not authenticated/authorized", method)
}
