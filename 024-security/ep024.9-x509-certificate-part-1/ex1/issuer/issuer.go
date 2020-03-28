package issuer

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"sec/ex8/serialize"
	"time"

	"github.com/sirupsen/logrus"
)

type (
	// Name contains information that to uniquely identifies a host
	Name struct {
		CommonName                       string
		OtherNames                       []string
		Street                           []string
		City, Locale, State, Province    []string
		Coutry                           []string
		Organization, OrganizationalUnit []string
		IPAddress                        []string
		Name                             string
	}
	Certificate struct {
		Subject             Name
		Issuer              Name
		Signature           []byte
		SubjectPublicKey    *rsa.PublicKey
		IsuerPublicKey      *rsa.PublicKey
		NotBefore, NotAfter time.Time
	}
)

var (
	rng           = rand.Reader
	issuerPrivKey = createKey()
	issuerName    = "Omni Trust"
)

func PublicKey() *rsa.PublicKey {
	return &issuerPrivKey.PublicKey
}

func NewSignedHostID(hostname string, ipAddrs []string) (*SignedHostID, error) {
	if len(hostname) < 2 {
		return nil, fmt.Errorf("Invalid hostname, hostname too short")
	}

	if len(ipAddrs) == 0 {
		return nil, fmt.Errorf("IpAddrs field must not be empty")
	}

	id := HostID{Name: hostname, IPAddress: ipAddrs}
	signedHostID, _ := signMyID(issuerPrivKey, id)

	return signedHostID, nil
}

func signMyID(key *rsa.PrivateKey, subject HostID) (*SignedHostID, error) {
	sigValue := SignedHostID{
		Subject: subject,
	}

	b := serialize.GobEncode(subject)
	h := sha256.Sum256(b)
	buf, err := rsa.SignPKCS1v15(rng, key, crypto.SHA256, h[:])

	if err != nil {
		return nil, fmt.Errorf("Unable to encrypt HostID hash: %v", err)
	}

	sigValue.Signature = buf

	issuer := Issuer{Name: issuerName, PublicKey: &key.PublicKey}
	sigValue.Issuer = issuer

	return &sigValue, nil
}

func createKey() *rsa.PrivateKey {
	const keySize = 1024 * 2
	k, err := rsa.GenerateKey(rng, keySize)
	if err != nil {
		logrus.Error("Unable to create host key:", err)
	}

	return k
}
