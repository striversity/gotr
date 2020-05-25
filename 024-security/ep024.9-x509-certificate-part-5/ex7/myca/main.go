package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

func main() {
	caFile := "myca"
	flag.StringVar(&caFile, "ca-filename", caFile, "file to write CA cert and private key")
	var clientName string
	flag.StringVar(&clientName, "client", "", "client hostname")
	flag.Parse()

	err := initRootCA(caFile)
	if err != nil {
		logrus.Fatalf("unable to init Root CA info: %v", err)
	}

	err = createClientCA(clientName)
	if err != nil {
		logrus.Error(err)
	}
}
