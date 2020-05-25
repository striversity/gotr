package main

import (
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
)

func getKeyID(key *rsa.PublicKey) []byte {
	buf := x509.MarshalPKCS1PublicKey(key)
	id := sha1.Sum(buf)
	return id[:]
}
