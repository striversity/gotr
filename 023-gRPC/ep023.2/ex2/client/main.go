package main

import (
	"context"
	"fmt"
	"log"
	"mms/model"

	"google.golang.org/grpc"
)

var (
	server = "localhost:8080"
)

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(server, opts...)
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to connect to gRPC service: %v", err))
	}
	defer conn.Close()

	client := model.NewMyMathServiceClient(conn)

	// call Add on client stub
	ctx := context.Background()
	in := &model.MathRequest{Operand1: 11, Operand2: 4}
	result, err := client.Add(ctx, in)
	if err != nil {
		log.Fatal(fmt.Errorf("Add rpc failed: %v", err))
	}

	fmt.Printf("Add(%v) => %v\n", in, result)
}
