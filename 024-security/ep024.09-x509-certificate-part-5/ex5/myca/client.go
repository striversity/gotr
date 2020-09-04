package main

import (
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"time"
)

func createClientCA(cn string) error {
	if len(cn) == 0 {
		return fmt.Errorf("invalid client name '%v'", cn)
	}

	clientInfo := &certInfo{}
	var err error

	fn := cn + "-key.pem"

	clientInfo.key, err = createKey(fn)
	if err != nil {
		return fmt.Errorf("uanble to create private key for client %v", cn)
	}

	err = initClientCert(cn, &clientInfo.key.PublicKey)
	return err
}

func initClientCert(cn string, key *rsa.PublicKey) error {
	err := createSignedCert(cn, key)
	if err != nil {
		return fmt.Errorf("unable to create cert for client %v: %v", cn, err)
	}

	return err
}

func createSignedCert(cn string, clientKey *rsa.PublicKey) error {
	client := &x509.Certificate{}
	client.SerialNumber = big.NewInt(time.Now().Unix())

	now := time.Now()
	client.NotBefore = now
	client.NotAfter = now.AddDate(1, 0, 0)

	client.Subject = pkix.Name{
		CommonName: cn,
	}

	client.SubjectKeyId = getKeyID(clientKey)

	der, err := x509.CreateCertificate(rng, client, rootCAInfo.cert, clientKey, rootCAInfo.key)
	if err != nil {
		return fmt.Errorf("unable to create cert: %v", err)
	}

	fn := cn + ".pem"
	err = exportCert(fn, der)
	if err != nil {
		return err
	}

	return err
}
