package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

func main() {
	fn := "id_rsa"
	flag.StringVar(&fn, "f", fn, "base filename for private and public key files")
	flag.Parse()

	key, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		logrus.Fatalf("uanble to generate SSH keys: %v", err)
	}

	// write keys to files
	err = saveKeys(fn, key)
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("Private and public keys created in %v and %v.pub respectively\n", fn, fn)
}

func saveKeys(fn string, key *rsa.PrivateKey) error {
	err := savePrivateKey(fn, key)
	if err != nil {
		return err
	}

	err = savePublicKey(fn+".pub", &key.PublicKey)
	if err != nil {
		os.Remove(fn) // remove private key file if we can't create public keey
		return err
	}

	return nil
}

func savePrivateKey(fn string, key *rsa.PrivateKey) error {
	f, err := os.Create(fn)
	if err != nil {
		return fmt.Errorf("unable to create SSH priavte key file: %v", err)
	}
	defer f.Close()

	buf := x509.MarshalPKCS1PrivateKey(key)
	b := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: buf,
	}
	return pem.Encode(f, b)
}

func savePublicKey(fn string, key *rsa.PublicKey) error {
	pubKey, err := ssh.NewPublicKey(key)
	if err != nil {
		return fmt.Errorf("unable to create SSH authtorized public key: %v", err)
	}

	buf := ssh.MarshalAuthorizedKey(pubKey)

	err = ioutil.WriteFile(fn, buf, 0644)
	if err != nil {
		return fmt.Errorf("unable to create SSH public key file: %v", err)
	}

	return nil
}
