package main

import (
	"benc/model"
	"bytes"
	"io"
	"math/rand"
	"net/http"
	"time"

	proto "github.com/golang/protobuf/proto"
	log "github.com/mgutz/logxi/v1"
)

const (
	requestsPerClient = 100000
	maxBatchSize      = (requestsPerClient / 10) * 2 // 20% of total request
)

var (
	s      = rand.NewSource(time.Now().Unix())
	r      = rand.New(s)
	logger = log.New("client")
)

func submitRequests(url string) {
	var req *model.ClientReq
	msgLeft := requestsPerClient
	var reqID uint64

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
			req.Data = make([]byte, 1024)
			buf := encodeReq(req)
			// fmt.Println(buf) // send to server
			resp, err := http.Post(url, "application/binary", buf)
			if nil != err {
				logger.Error("Post error:", err)
				break // try again later
			}
			defer resp.Body.Close()
		}
		// // pause a bit between batches
		// time.Sleep(time.Duration(r.Intn(200)) * time.Millisecond)
	}
}

func encodeReq(req *model.ClientReq) io.Reader {
	var buf = &bytes.Buffer{}
	b, err := proto.Marshal(req)
	if err != nil {
		log.Error("Unable to encode clent message in Protobuf", err)
		return buf
	}

	buf.Write(b)
	return buf
}
