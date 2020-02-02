package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

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
	v := sha1.Sum([]byte(s))
	return v[:]
}
