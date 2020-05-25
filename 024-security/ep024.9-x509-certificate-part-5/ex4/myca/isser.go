package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

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
	fn += ".pem"

	var buf []byte
	var err error
	var der []byte

	buf, err = ioutil.ReadFile(fn)
	if err != nil {
		logrus.Infof("no cert found in %v for root CA, creating", fn)

		der, err = createSelfSignedCert(fn, ci.key)
		if err != nil {
			return fmt.Errorf("unable to create root CA cert: %v", err)
		}
	} else {
		b, _ := pem.Decode(buf)
		der = b.Bytes
	}

	ci.cert, err = x509.ParseCertificate(der)
	return err
}

func createSelfSignedCert(fn string, key *rsa.PrivateKey) ([]byte, error) {
	issuer := &x509.Certificate{}
	issuer.SerialNumber = big.NewInt(time.Now().Unix())

	now := time.Now()
	issuer.NotBefore = now
	issuer.NotAfter = now.AddDate(10, 0, 0)

	issuer.Subject = pkix.Name{
		CommonName: "Omni Trust, Inc",
	}

	der, err := x509.CreateCertificate(rng, issuer, issuer, &key.PublicKey, key)
	if err != nil {
		return nil, fmt.Errorf("unable to create cert: %v", err)
	}

	err = exportCert(fn, der)
	if err != nil {
		return nil, err
	}

	return der, err
}

func exportCert(fn string, der []byte) error {
	b := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: der,
	}

	file, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("unable to create file %v: %v", fn, err)
	}
	defer file.Close()

	err = pem.Encode(file, b)
	if err != nil {
		return fmt.Errorf("unable encode certificate in PEM format: %v", err)
	}

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
