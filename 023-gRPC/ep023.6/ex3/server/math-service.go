package main

import (
	"context"
	"fmt"
	"log"
	"mms/auth"
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

	if ok, err := authorized(ctx, "Mod()"); !ok {
		return nil, err
	}

	if ok, err := authorized2(ctx, "Mod()"); !ok {
		return nil, err
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

	if ok, err := authorized2(ctx, "Mod()"); !ok {
		return nil, err
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
