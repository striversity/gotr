package main

import (
	"fmt"
	"math/rand"
	"thread-pool/model"
	"time"
)

const (
	requestsPerClient = 10000
	maxBatchSize      = (requestsPerClient / 10) * 2 // 20% of total request
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func main() {
	var req *model.ClientReq
	msgLeft := requestsPerClient
	var reqID uint

	for 0 < msgLeft {
		batch := r.Intn(maxBatchSize)
		if batch > msgLeft {
			batch = msgLeft
		}
		msgLeft -= batch

		for i := 0; i < batch; i++ {
			req = &model.ClientReq{}
			reqID++
			req.ID = reqID
			req.Size = r.Intn(model.ReqDataSize)
			for y := 0; y < req.Size; y++ {
				req.Data[y] = byte(y + 1)
			}
			fmt.Println(req) // send to server
		}
		// pause a bit between batches
		time.Sleep(time.Duration(r.Intn(200)) * time.Millisecond)
	}
}
