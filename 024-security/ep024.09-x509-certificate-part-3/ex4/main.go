package main

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sec/issuer"

	"github.com/sirupsen/logrus"
)

var (
	certOutFilename string
	formatPem       = false
	certInFilename  string
)

func main() {
	flag.StringVar(&certOutFilename, "out", certOutFilename, "Out cert filename")
	flag.BoolVar(&formatPem, "pem", formatPem, "Write cert in PEM(text) format, default is DER(binary)")
	flag.StringVar(&certInFilename, "in", certInFilename, "In cert filename")
	flag.Parse()

	if len(certOutFilename) == 0 && len(certInFilename) == 0 {
		logrus.Fatalf("At least 'in' or 'out' MUST be specified")
	}
	if len(certOutFilename) > 0 && len(certInFilename) > 0 {
		logrus.Fatalf("At least 'in' or 'out', but not both")
	}

	if len(certOutFilename) > 0 {
		createCert()
		return
	}

	readCert()
}

func readCert() {
	buf, err := ioutil.ReadFile(certInFilename)
	if err != nil {
		logrus.Fatalf("Unable to read file: %v", err)
	}

	// try to decode bytes as PEM
	block, _ := pem.Decode(buf)

	var cert *x509.Certificate
	if block != nil {
		cert, err = x509.ParseCertificate(block.Bytes)
	} else {
		cert, err = x509.ParseCertificate(buf)
	}

	if err != nil {
		logrus.Fatalf("Failed to parse certificate: %v", err)
	}

	fmt.Println("Successfully parsed certificate.")

	_ = cert
}

func createCert() {
	cert, err := issuer.NewSelfSignedCert()
	if err != nil {
		logrus.Fatalf("Unable to create certificate: %v", err)
	}

	var buf []byte
	if !formatPem {
		certOutFilename += ".der"
		buf = cert
	} else {
		certOutFilename += ".pem"

		block := &pem.Block{
			Type: "CERTIFICATE",
			Bytes: cert,
		}

		b := &bytes.Buffer{}
		err = pem.Encode(b, block)
		if err != nil {
			logrus.Fatalf("Unable encode certificate in PEM: %v", err)
		}

		buf = b.Bytes()
	}

	err = ioutil.WriteFile(certOutFilename, buf, 0644)
	if err != nil {
		logrus.Fatalf("Unable to write certificate: %v", err)
	}
}
