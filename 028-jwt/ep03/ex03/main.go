package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

var (
	sampleJWT = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEwNCwibmFtZSI6IkphbmUgRG9lIiwicm9sZXMiOlsiYWRtaW4iLCJ1c2VyIl0sImVtYWlsIjoiamFuZS5kb2VAZXhhbXBsZS5jb20iLCJpYXQiOjE1Nzc4MzY4MDB9.GzM3CiKknmSMjBtjEi74BuLJny5akDw1D8QhjjMIsTo`
	sections  = []string{"Header", "Payload", "Signature"}
)

func main() {
	parts := strings.Split(sampleJWT, ".")

	for i, part := range parts[:2] {
		fmt.Printf("%s\n", sections[i])
		fmt.Printf("\tEncoded: %s\n", part)
		decPart, _ := base64.StdEncoding.DecodeString(part)
		fmt.Printf("\tDecoded: %s\n", decPart)
	}
}
