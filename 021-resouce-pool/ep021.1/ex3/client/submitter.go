package main

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"thread-pool/model"
	"time"
)

const (
	requestsPerClient = 100000
	maxBatchSize      = (requestsPerClient / 10) * 2 // 20% of total request
)

var (
	s = rand.NewSource(time.Now().Unix())
	r = rand.New(s)
)

func submitRequests(url string) {
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
			buf := encodeReq(req)
			// fmt.Println(buf) // send to server
			resp, err := http.Post(url, "text/json", buf)
			if nil == err {
				defer resp.Body.Close()
			}
		}
		// pause a bit between batches
		time.Sleep(time.Duration(r.Intn(200)) * time.Millisecond)
	}
}

func encodeReq(req *model.ClientReq) io.Reader {
	var buf = &bytes.Buffer{}
	jsonEnc := json.NewEncoder(buf)
	jsonEnc.Encode(req)
	return buf
}
