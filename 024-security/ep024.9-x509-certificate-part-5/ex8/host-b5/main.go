package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	var caCertFilename string
	flag.StringVar(&caCertFilename, "cacert", caCertFilename, "filename containing ca cert")
	var insecure bool
	flag.BoolVar(&insecure, "k", false, "Accept/Ignore all server SSL certificates")
	flag.Parse()

	url := flag.Arg(0)

	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	caCertPem, err := ioutil.ReadFile(caCertFilename)
	if err == nil {
		// append our to the cert pool
		rootCAs.AppendCertsFromPEM(caCertPem)
	}

	config := &tls.Config{
		InsecureSkipVerify: insecure,
		RootCAs:            rootCAs,
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}

	client := &http.Client{
		Transport: tr,
	}

	resp, err := client.Get(url)

	if err != nil {
		logrus.Fatalf("unable to connect to server %v: %v", url, err)
	}

	body := resp.Body
	defer body.Close()

	io.Copy(os.Stdout, resp.Body)
	fmt.Println()
}
