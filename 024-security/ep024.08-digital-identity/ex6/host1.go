package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

type (
	SignedHostID struct {
		Subject   HostID         // Subject information, basically who is this signed ID for?
		Signature []byte         // signed hash of HostID data serialized
		PublicKey *rsa.PublicKey // corresponding pub key of private used for signature
	}
)

func host1(out chan<- []byte, in <-chan []byte) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		myID := HostID{Name: "host1", IPAddress: []string{"10.10.1.2"}}
		myKey := createKey()

		hostSig, err := signMyID(myKey, myID)

		fmt.Println("signMyID():", err)
		out <- gobEncode(hostSig)
		msg := <-in
		fmt.Println("Host2 responded with:", string(msg))
	}()

}

func signMyID(key *rsa.PrivateKey, id HostID) (*SignedHostID, error) {
	sigValue := SignedHostID{
		Subject: id, PublicKey: &key.PublicKey,
	}

	b := gobEncode(id)
	h := sha256.Sum256(b)
	buf, err := rsa.SignPKCS1v15(rng, key, crypto.SHA256, h[:])

	if err != nil {
		return nil, fmt.Errorf("Unable to encrypt HostID hash: %v", err)
	}

	sigValue.Signature = buf
	return &sigValue, nil
}
