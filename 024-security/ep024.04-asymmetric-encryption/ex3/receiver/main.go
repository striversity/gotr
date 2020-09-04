package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	keyFile = "key.priv"
	rnd     = rand.Reader
)

func main() {
	flag.StringVar(&keyFile, "k", keyFile, "Private key file name")
	flag.Parse()

	// read in the private key
	key := &rsa.PrivateKey{}
	f, _ := os.Open(keyFile)
	w := gob.NewDecoder(f)
	err := w.Decode(key)
	if err != nil {
		log.Fatalf("Unable to read key from %v: %v", keyFile, err)
	}
	f.Close()

	// read message from stdin until end and write to stdout
	msg, _ := ioutil.ReadAll(os.Stdin)

	// encrypt message to owner of keypair
	buf, err := key.Decrypt(rnd, msg, nil)
	if err != nil {
		log.Fatalf("Unable to decrypt message: %v", err)
	}

	fmt.Println(string(buf))
}
