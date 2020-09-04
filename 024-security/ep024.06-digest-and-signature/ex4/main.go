package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

var header = []byte("This is my header")

func main() {
	for _, v := range os.Args[1:] {
		h := hash(v)
		encodedStr := hex.EncodeToString(h)
		fmt.Println(v, encodedStr)
	}
}
func hash(fn string) []byte {
	f, err := os.Open(fn)
	if err != nil {
		return nil
	}
	defer f.Close()

	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return nil
	}

	return h.Sum(nil)
}
