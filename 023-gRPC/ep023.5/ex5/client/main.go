package main

import (
	"benc/model"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var a int64 = 5
	var b int64 = 6
	req := &model.BiOperReq{OperandA: a, OperandB: b}
	fmt.Printf("Info - req: %v\n", req)

	// send requrest to gRPC server
	// ..
	conn, err := grpc.Dial("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := model.NewMathOperServiceClient(conn)
	resp := client.Add(context.Background(), req)
	fmt.Printf("Info - req: %v\n", resp)
}
