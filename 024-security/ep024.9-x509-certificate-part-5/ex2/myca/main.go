package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

func main() {
	caFile := "myca"
	flag.StringVar(&caFile, "ca-filename", caFile, "file to write CA cert and private key")
	flag.Parse()

	err := initRootCA(caFile)
	if err != nil {
		logrus.Fatalf("unable to init Root CA info: %v", err)
	}
}
