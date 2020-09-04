package main

import (
	"flag"
	"io/ioutil"
	"sec/issuer"

	"github.com/sirupsen/logrus"
)

var (
	caCertFilename = "ca_cert.der"
)

func main() {
	flag.StringVar(&caCertFilename, "out", caCertFilename, "Out cert filename")
	flag.Parse()

	cert, err := issuer.NewSelfSignedCert()
	if err != nil {
		logrus.Fatalf("Unable to create certificate: %v", err)
	}

	err = ioutil.WriteFile(caCertFilename, cert, 0644)
	if err != nil {
		logrus.Fatalf("Unable to write certificate: %v", err)
	}
}
