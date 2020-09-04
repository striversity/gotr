package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

var header = []byte("This is my header")

func main() {
	var v string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		v = scanner.Text()
		h := hash(v)
		encodedStr := hex.EncodeToString(h)
		fmt.Println(v, encodedStr)
	}
}
func hash(s string) []byte {
	h := sha1.New()
	h.Write(header)
	h.Write([]byte(s))
	v := h.Sum(nil)
	return v[:]
}
