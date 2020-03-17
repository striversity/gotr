package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"sec/ex8/issuer"
	"sec/ex8/serialize"

	"github.com/sirupsen/logrus"
)

func host2(out chan<- []byte, in <-chan []byte) {
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

		if !isTrustedHost(req.From) {
			m := "ALARM: Untrusted host"
			logrus.Errorf(m)
			out <- []byte(m)
			return
		}

		reply := fmt.Sprintf("Hi %v, here is the SECRET", req.From.Subject.Name)
		out <- []byte(reply)
	}()

}

func isTrustedHost(identity *issuer.SignedHostID) bool {
	buf := serialize.GobEncode(identity.Subject)
	h := sha256.Sum256(buf)

	clientPubKey := identity.Issuer.PublicKey
	err := rsa.VerifyPKCS1v15(clientPubKey, crypto.SHA256, h[:], identity.Signature)

	if err != nil {
		return false
	}

	logrus.Errorf("Signature for %v is valid, but host signer unkown", identity.Subject.Name)

	return false
}
