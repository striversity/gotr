package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Please enter secure message, end with CTRL+D:")
	msg, _ := ioutil.ReadAll(os.Stdin)

	key := []byte("simple password")
	sharedKey, _ := bcrypt.GenerateFromPassword(key, bcrypt.DefaultCost)
	sharedKey = sharedKey[:32]
	fmt.Println("sharedKey:", len(sharedKey), sharedKey)

	buf := encrypt(msg, sharedKey)
	fmt.Println("Cyptered Output:", hex.EncodeToString(buf))

	fmt.Println("--------------------")
	buf = decrypt(buf, sharedKey)
	fmt.Println("Recovered Output:", string(buf))
}

func decrypt(buf, key []byte) []byte {
	block, err := aes.NewCipher(key)
	fmt.Println("NewCipher error:", err)
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewOFB(block, iv)

	in := bytes.NewReader(buf)
	reader := &cipher.StreamReader{S: stream, R: in}

	out := &bytes.Buffer{}
	io.Copy(out, reader)
	return out.Bytes()
}

func encrypt(text, key []byte) []byte {
	block, err := aes.NewCipher(key)
	fmt.Println("newCipher error:", err)
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewOFB(block, iv)

	out := &bytes.Buffer{}
	writer := &cipher.StreamWriter{S: stream, W: out}
	writer.Write(text)

	return out.Bytes()
}
