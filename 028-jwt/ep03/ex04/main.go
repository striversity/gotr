package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

var (
	sampleJWT = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEwNCwibmFtZSI6IkphbmUgRG9lIiwicm9sZXMiOlsiYWRtaW4iLCJ1c2VyIl0sImVtYWlsIjoiamFuZS5kb2VAZXhhbXBsZS5jb20iLCJpYXQiOjE1Nzc4MzY4MDB9.GzM3CiKknmSMjBtjEi74BuLJny5akDw1D8QhjjMIsTo`
	sharedKey = "very-secure"
)

func main() {
	parts := strings.Split(sampleJWT, ".")
	header := parts[0]
	payload := parts[1]
	signature := parts[2]

	data := header + "." + payload

	fmt.Printf("Signature is valid: %t\n", verifySignature(data, signature, sharedKey))
}

func verifySignature(data, signature, sharedKey string) bool {
	// decode the signature from base64
	decodedSig, err := base64.RawURLEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println("Failed to decode signature: ", err)
		return false
	}

	// create a new HMAC-SHA256 hasher wit the secret
	h := hmac.New(sha256.New, []byte(sharedKey))

	// write the data to the hasher
	_, err = h.Write([]byte(data))
	if err != nil {
		fmt.Println("Failed to write to hasher: ", err)
		return false
	}

	// calculate checksum
	calculatedSig := h.Sum(nil)

	// compare the signatures
	return hmac.Equal(decodedSig, calculatedSig)
}
