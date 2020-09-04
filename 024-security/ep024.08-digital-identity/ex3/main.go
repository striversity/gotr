package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)

type (
	// HostID contains information that to uniquely identifies a host
	HostID struct {
		Name      string
		IPAddress []string
	}
)

var (
	wg sync.WaitGroup
)

func main() {
	ch1 := make(chan []byte)
	ch2 := make(chan []byte)

	host1(ch1, ch2)
	host2(ch2, ch1)

	// wait for gorountines
	wg.Wait()
}

func gobDecode(v interface{}, buf []byte) error {
	if v == nil || buf == nil {
		return fmt.Errorf("Invalid paraemters, v nor buf cannot be nil")
	}

	b := bytes.NewReader(buf)
	dec := gob.NewDecoder(b)

	return dec.Decode(v)
}

func gobEncode(v interface{}) []byte {
	b := &bytes.Buffer{}
	enc := gob.NewEncoder(b)
	enc.Encode(v)
	return b.Bytes()
}
