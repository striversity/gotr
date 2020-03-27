package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	server = "localhost:8080"
)

func main() {
	flag.StringVar(&server, "s", server, "gRPC server server host:port")
	flag.Parse()

	url := fmt.Sprintf("https://%v/hello", server)

	certPool := x509.NewCertPool()
	cert, _ := ioutil.ReadFile("../cert.pem")
	certPool.AppendCertsFromPEM(cert)
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}
	tr := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Unable to connect to %v\n", server)
	}

	defer resp.Body.Close()

	fmt.Printf("Response from %s:\n", server)
	io.Copy(os.Stdout, resp.Body)
	fmt.Println()
}
