package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

type (
	// HostID contains information that to uniquely identifies a host
	HostID struct {
		Name      string
		IPAddress []string
	}

	SignedHostID struct {
		Subject   HostID         // Subject information, basically who is this signed ID for?
		Signature []byte         // signed hash of HostID data serialized
		PublicKey *rsa.PublicKey // corresponding pub key of private used for signature
	}
)

var (
	wg  sync.WaitGroup
	rng = rand.Reader
)

func main() {
	ch1 := make(chan []byte)
	ch2 := make(chan []byte)

	host1(ch1, ch2)
	host2(ch2, ch1)

	// wait for gorountines
	wg.Wait()
}
func createKey() *rsa.PrivateKey {
	const keySize = 1024 * 2
	k, err := rsa.GenerateKey(rng, keySize)
	if err != nil {
		logrus.Error("Unable to create host key:", err)
	}

	return k
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
