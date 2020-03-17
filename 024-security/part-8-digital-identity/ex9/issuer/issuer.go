package issuer

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"sec/ex8/serialize"

	"github.com/sirupsen/logrus"
)

type (
	// HostID contains information that to uniquely identifies a host
	HostID struct {
		Name      string
		IPAddress []string
	}

	Issuer struct {
		Name      string
		PublicKey *rsa.PublicKey //  public key of common trust, signer of identity
	}
	SignedHostID struct {
		Subject   HostID // Subject information, basically who is this signed ID for?
		Issuer    Issuer
		Signature []byte // signed hash of HostID data serialized
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
