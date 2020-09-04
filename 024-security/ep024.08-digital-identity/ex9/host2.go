package main

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"sec/ex9/issuer"
	"sec/ex9/serialize"

	"github.com/sirupsen/logrus"
)

func host2(out chan<- []byte, in <-chan []byte, trustedKey *rsa.PublicKey) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		msg := <-in
		var req Request
		err := serialize.GobDecode(&req, msg)
		if err != nil {
			logrus.Errorf("Unable to decode client message: %v", err)
			out <- []byte("ERROR: try again")
			return
		}

		if !isTrustedHost(req.From, trustedKey) {
			m := "ALARM: Untrusted host"
			logrus.Errorf(m)
			out <- []byte(m)
			return
		}

		reply := fmt.Sprintf("Hi %v, here is the SECRET", req.From.Subject.Name)
		out <- []byte(reply)
	}()

}

func isTrustedHost(identity *issuer.SignedHostID, trustedKey *rsa.PublicKey) bool {
	buf := serialize.GobEncode(identity.Subject)
	h := sha256.Sum256(buf)

	signerKey := identity.Issuer.PublicKey
	err := rsa.VerifyPKCS1v15(signerKey, crypto.SHA256, h[:], identity.Signature)

	if err != nil {
		return false
	}

	b1 := serialize.GobEncode(signerKey)
	b2 := serialize.GobEncode(trustedKey)

	if !bytes.Equal(b1, b2) {
		logrus.Errorf("Signature valid, but not a trusted issuer: %v", identity.Issuer.Name)
		return false
	}

	return true
}
