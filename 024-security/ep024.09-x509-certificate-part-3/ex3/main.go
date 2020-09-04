package main

import (
	"bytes"
	"encoding/pem"
	"flag"
	"io/ioutil"
	"os"
	"sec/issuer"

	"github.com/sirupsen/logrus"
)

var (
	caCertFilename = "ca_cert"
	formatPem      = false
)

func main() {
	flag.StringVar(&caCertFilename, "out", caCertFilename, "Out cert filename")
	flag.BoolVar(&formatPem, "pem", formatPem, "Write cert in PEM(text) format, default is DER(binary)")
	flag.Parse()

	cert, err := issuer.NewSelfSignedCert()
	if err != nil {
		logrus.Fatalf("Unable to create certificate: %v", err)
	}

	var buf []byte
	if !formatPem {
		caCertFilename += ".der"
		buf = cert
	} else {
		caCertFilename += ".pem"

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

	err = ioutil.WriteFile(caCertFilename, buf, 0644)
	if err != nil {
		logrus.Fatalf("Unable to write certificate: %v", err)
	}
}
