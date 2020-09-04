package main

import (
	"crypto/rsa"
	"crypto/x509"
)

type certInfo struct {
	key  *rsa.PrivateKey
	cert *x509.Certificate
}

var (
	rootCAInfo *certInfo
)

func initRootCA(fn string) error {
	rootCAInfo = &certInfo{}
	rootCAInfo.initKey(fn)
	rootCAInfo.initRootCert(fn)
	return nil
}

func (ci *certInfo) initKey(fn string) error {
	return nil
}

func (ci *certInfo) initRootCert(fn string) error {
	return nil
}
