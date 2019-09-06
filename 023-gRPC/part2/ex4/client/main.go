package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mms/model"

	"google.golang.org/grpc"
)

var (
	server = "localhost:8080"
)

func main() {
	flag.StringVar(&server, "s", server, "gRPC server address host:port")
	flag.Parse()

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(server, opts...)
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to connect to gRPC service: %v", err))
	}
	defer conn.Close()

	client := model.NewMyMathServiceClient(conn)
	testMathService(client)
}

func testMathService(client model.MyMathServiceClient) {

	ctx := context.Background()
	in := &model.MathRequest{Operand1: 11, Operand2: 4}

	// call Add on client stub
	result, err := client.Add(ctx, in)
	if err != nil {
		log.Fatal(fmt.Errorf("Add rpc failed: %v", err))
	}

	fmt.Printf("Add(%v) => %v\n", in, result)

	// call Sub on client stub
	result, err = client.Sub(ctx, in)
	if err != nil {
		log.Fatal(fmt.Errorf("Sub rpc failed: %v", err))
	}

	fmt.Printf("Sub(%v) => %v\n", in, result)

	// call Mul on client stub
	result, err = client.Mul(ctx, in)
	if err != nil {
		log.Fatal(fmt.Errorf("Mul rpc failed: %v", err))
	}

	fmt.Printf("Mul(%v) => %v\n", in, result)

	// call Div on client stub
	result, err = client.Div(ctx, in)
	if err != nil {
		log.Fatal(fmt.Errorf("Div rpc failed: %v", err))
	}

	fmt.Printf("Div(%v) => %v\n", in, result)

	// call Mod on client stub
	result, err = client.Mod(ctx, in)
	if err != nil {
		log.Fatal(fmt.Errorf("Mod rpc failed: %v", err))
	}

	fmt.Printf("Mod(%v) => %v\n", in, result)
}
