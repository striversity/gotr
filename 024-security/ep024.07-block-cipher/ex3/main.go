package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type (
	Payload struct {
		Message []byte
		EncKey  []byte
	}
)

func main() {
	fmt.Println("Please enter secure message, end with CTRL+D:")
	msg, _ := ioutil.ReadAll(os.Stdin)

	key := []byte("simple password")

	buf := encrypt(msg, key)
	fmt.Println("Cyptered Output:", hex.EncodeToString(buf))

	fmt.Println("--------------------")
	buf = decrypt(buf)
	fmt.Println("Recovered Output:", string(buf))
}

func decrypt(buf []byte) []byte {
	b := bytes.NewReader(buf)
	dec := gob.NewDecoder(b)
	payload := &Payload{}
	err := dec.Decode(payload)
	fmt.Println("gob.Decode error:", err)

	sharedKey := payload.EncKey
	fmt.Println("sharedKey:", len(sharedKey), sharedKey)

	block, err := aes.NewCipher(sharedKey)
	fmt.Println("NewCipher error:", err)
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewOFB(block, iv)

	in := bytes.NewReader(payload.Message)
	reader := &cipher.StreamReader{S: stream, R: in}

	out := &bytes.Buffer{}
	io.Copy(out, reader)
	return out.Bytes()
}

func encrypt(text, key []byte) []byte {
	sharedKey, _ := bcrypt.GenerateFromPassword(key, bcrypt.DefaultCost)
	sharedKey = sharedKey[:32]
	fmt.Println("sharedKey:", len(sharedKey), sharedKey)

	block, err := aes.NewCipher(sharedKey)
	fmt.Println("newCipher error:", err)
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewOFB(block, iv)

	out := &bytes.Buffer{}
	writer := &cipher.StreamWriter{S: stream, W: out}
	writer.Write(text)

	payload := &Payload{Message: out.Bytes(), EncKey: sharedKey}

	b := &bytes.Buffer{}
	enc := gob.NewEncoder(b)
	enc.Encode(payload)

	return b.Bytes()
}
