package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var (
	keyFile = "key.pub"
	rnd     = rand.Reader
)

func main() {
	flag.StringVar(&keyFile, "k", keyFile, "Public key file name")
	flag.Parse()

	// read in the public key
	key := &rsa.PublicKey{}
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
	buf, err := rsa.EncryptPKCS1v15(rnd, key, msg)
	if err != nil {
		log.Fatalf("Unable to encrypt message: %v", err)
	}

	os.Stdout.Write(buf)
}
