package main

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

const (
	keySize = 1024 * 2
)

var (
	rng = rand.Reader
)

type (
	Payload struct {
		Message   []byte
		Signature []byte
	}
	Data struct {
		Payload []byte
		EncKey  []byte
	}
)

func main() {
	sendKey, _ := rsa.GenerateKey(rng, keySize)
	recvKey, _ := rsa.GenerateKey(rng, keySize)

	fmt.Println("Please enter secure message, end with CTRL+D:")
	msg, _ := ioutil.ReadAll(os.Stdin)

	buf := encrypt(msg, sendKey, &recvKey.PublicKey)
	fmt.Println("Cyptered Output:", hex.EncodeToString(buf))

	fmt.Println("--------------------")
	buf = decrypt(buf, recvKey, &sendKey.PublicKey)
	fmt.Println("Recovered Output:", string(buf))
}

func decrypt(buf []byte, recvKey *rsa.PrivateKey, sendKey *rsa.PublicKey) []byte {
	b := bytes.NewReader(buf)
	data := &Data{}
	dec := gob.NewDecoder(b)
	err := dec.Decode(data)
	fmt.Println("gob.Decode error:", err)

	sharedKey, err := rsa.DecryptPKCS1v15(rng, recvKey, data.EncKey)
	fmt.Println("sharedKey decrypt error:", err)

	fmt.Println("sharedKey:", len(sharedKey), sharedKey)

	block, err := aes.NewCipher(sharedKey[:32])
	fmt.Println("NewCipher error:", err)
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewOFB(block, iv)
	in := bytes.NewReader(data.Payload)
	reader := &cipher.StreamReader{S: stream, R: in}

	out := &bytes.Buffer{}
	io.Copy(out, reader)

	payload := &Payload{}
	dec = gob.NewDecoder(out)
	err = dec.Decode(payload)
	fmt.Println("gob.Decode error:", err)

	hash := sha256.Sum256(payload.Message)
	err = rsa.VerifyPKCS1v15(sendKey, crypto.SHA256, hash[:], payload.Signature)
	if err != nil {
		fmt.Println("BAD message signature:", err)
	} else {
		fmt.Println("Message signature valid")
	}
	return payload.Message
}

func encrypt(buf []byte, sendKey *rsa.PrivateKey, recvKey *rsa.PublicKey) []byte {
	hash := sha256.Sum256(buf)
	sig, _ := rsa.SignPKCS1v15(rng, sendKey, crypto.SHA256, hash[:])
	payload := &Payload{Message: buf, Signature: sig}
	b := &bytes.Buffer{}
	enc := gob.NewEncoder(b)
	enc.Encode(payload)

	key := []byte("simple password")
	sharedKey, _ := bcrypt.GenerateFromPassword(key, bcrypt.DefaultCost)
	sharedKey = sharedKey[:32]
	fmt.Println("sharedKey:", len(sharedKey), sharedKey)

	encKey, err := rsa.EncryptPKCS1v15(rng, recvKey, sharedKey)
	fmt.Println("encKey error:", err)

	out := &bytes.Buffer{}
	block, err := aes.NewCipher(sharedKey[:32])
	fmt.Println("NewCipher error:", err)
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewOFB(block, iv)
	writer := &cipher.StreamWriter{S: stream, W: out}
	io.Copy(writer, b)

	data := &Data{Payload: out.Bytes(), EncKey: encKey}
	b = &bytes.Buffer{}
	enc = gob.NewEncoder(b)
	enc.Encode(data)

	return b.Bytes()
}
