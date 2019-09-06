package main

import (
	"context"
	"fmt"
	"mms/model"
)

type (
	myMathService struct {
	}
)

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
