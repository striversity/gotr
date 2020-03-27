package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"log"
	"os"
)

const (
	keySize = 1024
)

func main() {
	rnd := rand.Reader
	priv, err := rsa.GenerateKey(rnd, keySize)
	if err != nil {
		log.Fatalf("Error generating key, unable to continue: %v", err)
	}

	// fmt.Printf("Private key: %v\n", priv)
	// fmt.Printf("Public key: %v\n", priv.Public())

	writePrivKeyFile(priv)
	writePubKeyFile(priv)
}
func writePrivKeyFile(priv *rsa.PrivateKey) {
	privKeyFile := "key.priv"
	f, _ := os.Create(privKeyFile)
	defer f.Close()
	w := gob.NewEncoder(f)
	w.Encode(priv)
}
func writePubKeyFile(priv *rsa.PrivateKey) {
	privKeyFile := "key.pub"
	f, _ := os.Create(privKeyFile)
	defer f.Close()
	w := gob.NewEncoder(f)
	w.Encode(priv.Public())
}
