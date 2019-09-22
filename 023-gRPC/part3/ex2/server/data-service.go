package main

import (
	"fmt"
	"math/rand"
	"mms/model"
	"time"
)

type (
	myDataService struct {
	}
)

var (
	src = rand.NewSource(time.Now().Unix())
	r   = rand.New(src)
)

func (m *myDataService) Random(in *model.RandomRequest, out model.DataService_RandomServer) error {
	if m == nil {
		return fmt.Errorf("Random called on nil object")
	}
	if in == nil {
		return fmt.Errorf("Random called with invalid paramter value nil")
	}

	count := int(in.Count)
	var v int64

	if in.Bounded {
		for i := 0; i < count; i++ {
			v = r.Int63n(in.MaxValue-in.MinValue) + in.MinValue
			out.Send(&model.RandomResponse{Value: v})
		}
		return nil
	}

	for i := 0; i < count; i++ {
		v = r.Int63()
		out.Send(&model.RandomResponse{Value: v})
	}

	return nil
}
