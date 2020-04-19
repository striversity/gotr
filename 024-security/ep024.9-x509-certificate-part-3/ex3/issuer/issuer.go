package issuer

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	rng           = rand.Reader
	issuerPrivKey = createKey()
	issuerName    = "Omni Trust"
	issuer        *x509.Certificate
)

func PublicKey() *rsa.PublicKey {
	return &issuerPrivKey.PublicKey
}

func NewSelfSignedCert() ([]byte, error) {
	issuer = &x509.Certificate{}
	issuer.SerialNumber = big.NewInt(time.Now().Unix())

	now := time.Now()
	issuer.NotBefore = now
	issuer.NotAfter = now.AddDate(1, 0, 0)

	issuer.Subject = isserInfo()

	issuer.DNSNames = []string{"vee-mbp.mv.lorrev.org"}
	issuer.EmailAddresses = []string{"vla@lorrev.org"}

	cert, err := x509.CreateCertificate(rng, issuer, issuer, &issuerPrivKey.PublicKey, issuerPrivKey)

	return cert, err
}

func isserInfo() pkix.Name {
	name := pkix.Name{
		CommonName:         issuerName,
		Country:            []string{"US"},
		Organization:       []string{"Striversity"},
		OrganizationalUnit: []string{"Media"},
		Province:           []string{"State"},
		Locality:           []string{"City"},
		StreetAddress:      []string{"123 First Street"},
	}

	return name
}
func createKey() *rsa.PrivateKey {
	const keySize = 1024 * 2
	k, err := rsa.GenerateKey(rng, keySize)
	if err != nil {
		logrus.Error("Unable to create host key:", err)
	}

	return k
}
