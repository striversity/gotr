package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
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

const (
	keySize = 1024 * 2
)

var (
	rng = rand.Reader
)

func main() {
	fmt.Println("Please enter secure message, end with CTRL+D:")
	msg, _ := ioutil.ReadAll(os.Stdin)

	recv, _ := rsa.GenerateKey(rng, keySize)
	attacker, _ := rsa.GenerateKey(rng, keySize)

	buf := encrypt(msg, &recv.PublicKey)
	fmt.Println("Cyptered Output:", hex.EncodeToString(buf))

	fmt.Println("--------------------")
	buf = decrypt(buf, recv)
	fmt.Println("Recovered Output:", string(buf))
	fmt.Println("ATTACKER: --------------------")
	buf = decrypt(buf, attacker)
	fmt.Println("Recovered Output:", string(buf))
}

func decrypt(buf []byte, recvKey *rsa.PrivateKey) []byte {
	b := bytes.NewReader(buf)
	dec := gob.NewDecoder(b)
	payload := &Payload{}
	err := dec.Decode(payload)
	fmt.Println("gob.Decode error:", err)

	sharedKey, err := rsa.DecryptPKCS1v15(rng, recvKey, payload.EncKey)
	fmt.Println("sharedKey decrypt error:", err)
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

func encrypt(text []byte, recvKey *rsa.PublicKey) []byte {
	key := []byte("simple password")
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

	sharedKey, err = rsa.EncryptPKCS1v15(rng, recvKey, sharedKey)
	fmt.Println("sharedKey encrypt error:", err)
	payload := &Payload{Message: out.Bytes(), EncKey: sharedKey}

	b := &bytes.Buffer{}
	enc := gob.NewEncoder(b)
	enc.Encode(payload)

	return b.Bytes()
}
