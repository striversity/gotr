package main

import (
	"crypto/rand"
	"crypto/rsa"
	"flag"
	"log"
	"sec/ex3/keyutil"
)

var (
	keySize     = 2048
	keyFilename = "rsa_id"
)

func main() {
	flag.IntVar(&keySize, "s", keySize, "Key size in bytes")
	flag.Parse()

	if len(flag.Args()) > 0 {
		keyFilename = flag.Arg(0)
	}

	rnd := rand.Reader
	priv, err := rsa.GenerateKey(rnd, keySize)
	if err != nil {
		log.Fatalf("Error generating key, unable to continue: %v", err)
	}

	keyutil.WritePrivKeyFile(keyFilename, priv)
	keyutil.WritePubKeyFile(keyFilename+".pub", priv)
}
