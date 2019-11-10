package main

import (
	"context"
	"fmt"
	"log"
	"mms/model"

	"google.golang.org/grpc/metadata"
)

type (
	myMathService struct {
	}
)

func (m *myMathService) Mod(ctx context.Context, in *model.MathRequest) (*model.MathResponse, error) {
	if m == nil {
		return nil, fmt.Errorf("Mod called on nil object")
	}
	if in == nil {
		return nil, fmt.Errorf("Mod called with invalid paramter value nil")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("Mod() - No metadata provided")
		return nil, fmt.Errorf("Mod() - No metadata provided")
	}

	log.Printf("Mod() - Metadata len: %v", md.Len())

	for k, v := range md {
		log.Printf("Mod() - Metadata key: %v -> %v", k, v)
	}

	if in.GetOperand2() == 0 {
		return nil, fmt.Errorf("Mod called with invalid paramter operand2 == 0")
	}

	result := &model.MathResponse{}
	result.Result = in.Operand1 % in.Operand2

	return result, nil
}

func (m *myMathService) Div(ctx context.Context, in *model.MathRequest) (*model.MathResponse, error) {
	if m == nil {
		return nil, fmt.Errorf("Div called on nil object")
	}
	if in == nil {
		return nil, fmt.Errorf("Div called with invalid paramter value nil")
	}
	if in.GetOperand2() == 0 {
		return nil, fmt.Errorf("Div called with invalid paramter operand2 == 0")
	}

	result := &model.MathResponse{}
	result.Result = in.Operand1 / in.Operand2

	return result, nil
}

func (m *myMathService) Mul(ctx context.Context, in *model.MathRequest) (*model.MathResponse, error) {
	if m == nil {
		return nil, fmt.Errorf("Mul called on nil object")
	}
	if in == nil {
		return nil, fmt.Errorf("Mul called with invalid paramter value nil")
	}

	result := &model.MathResponse{}
	result.Result = in.Operand1 * in.Operand2

	return result, nil
}

func (m *myMathService) Sub(ctx context.Context, in *model.MathRequest) (*model.MathResponse, error) {
	if m == nil {
		return nil, fmt.Errorf("Sub called on nil object")
	}
	if in == nil {
		return nil, fmt.Errorf("Sub called with invalid paramter value nil")
	}

	result := &model.MathResponse{}
	result.Result = in.Operand1 - in.Operand2

	return result, nil
}

func (m *myMathService) Add(ctx context.Context, in *model.MathRequest) (*model.MathResponse, error) {
	if m == nil {
		return nil, fmt.Errorf("Add called on nil object")
	}
	if in == nil {
		return nil, fmt.Errorf("Add called with invalid paramter value nil")
	}

	result := &model.MathResponse{}
	result.Result = in.Operand1 + in.Operand2

	return result, nil
}
