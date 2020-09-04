package issuer

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	rng            = rand.Reader
	issuerPrivKey  = createKey()
	issuerName     = "Omni Trust"
	issuer         *x509.Certificate
	caCertFilename = "ca_cert.der"
)

func PublicKey() *rsa.PublicKey {
	return &issuerPrivKey.PublicKey
}

func NewSelfSignedCert() error {
	issuer = &x509.Certificate{}
	issuer.SerialNumber = big.NewInt(time.Now().Unix())

	cert, err := x509.CreateCertificate(rng, issuer, issuer, &issuerPrivKey.PublicKey, issuerPrivKey)
	if err != nil {
		logrus.Errorf("Unable to create certificate: %v", err)
		return err
	}

	err = ioutil.WriteFile(caCertFilename, cert, 0644)
	if err != nil {
		logrus.Errorf("Unable to write certificate: %v", err)
		return err
	}

	return nil
}

func createKey() *rsa.PrivateKey {
	const keySize = 1024 * 2
	k, err := rsa.GenerateKey(rng, keySize)
	if err != nil {
		logrus.Error("Unable to create host key:", err)
	}

	return k
}
