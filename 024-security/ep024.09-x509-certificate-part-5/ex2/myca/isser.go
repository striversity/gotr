package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

type certInfo struct {
	key  *rsa.PrivateKey
	cert *x509.Certificate
}

var (
	rng        = rand.Reader
	rootCAInfo *certInfo
)

func initRootCA(fn string) error {
	rootCAInfo = &certInfo{}
	err := rootCAInfo.initKey(fn)
	if err != nil {
		return fmt.Errorf("unable to init root CA private key: %v", err)
	}

	err = rootCAInfo.initRootCert(fn)
	return err
}

// initKey initializes the field certInfo.key with rsa.Private key in the file
// fn + "-key.pem" or creates a new key
func (ci *certInfo) initKey(fn string) error {
	fn += "-key.pem"
	// try to read private key for Root CA from file
	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		logrus.Infof("no exiting key found in %v for root ca, creating", fn)
		// if we can't read the file, assume first run and generate a new private key
		ci.key, err = createKey(fn)
		return err
	}

	// we have a key file, so decode PEM bytes to rsa.PrivateKey
	b, _ := pem.Decode(buf)
	ci.key, err = x509.ParsePKCS1PrivateKey(b.Bytes)
	return err
}

// initRootCert initializes the certInfo.cert field with a cert saved in fn+".pem"
// or creates a new root cert (self-signed cert)
func (ci *certInfo) initRootCert(fn string) error {
	return nil
}

// createKey creates a rsa.PrivateKey and saves it to a file fn before returning the key
func createKey(fn string) (*rsa.PrivateKey, error) {
	const keySize = 1024 * 2
	k, err := rsa.GenerateKey(rng, keySize)
	if err != nil {
		return nil, fmt.Errorf("unable to create private key for %v: %v", fn, err)
	}

	b := &pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(k),
		Type:  "RSA PRIVATE KEY",
	}

	buf := &bytes.Buffer{}
	err = pem.Encode(buf, b)
	if err != nil {
		return nil, fmt.Errorf("Unable to encode private key to PEM: %v", err)
	}

	err = ioutil.WriteFile(fn, buf.Bytes(), 0600)
	return k, err
}
