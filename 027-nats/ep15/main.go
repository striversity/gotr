package main

import (
	"flag"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/nats-io/nats.go"
)

const putFileHelp = `put <store> <file>`
const getFileHelp = `get <store> <file> <destination>`

var (
	username string
	password string
	hostname = "localhost"
	port     = 4222
)

func init() {
	flag.StringVar(&username, "u", username, "username for NATS Server")
	flag.StringVar(&password, "p", password, "password for NATS Server")
	flag.StringVar(&hostname, "host", hostname, "NATS Server hostname")
	flag.IntVar(&port, "port", port, "NATS Server port")
	flag.Parse()
}

func main() {
	if len(flag.Args()) == 0 {
		log.Errorf("missing arguments, use either %s or %s", putFileHelp, getFileHelp)
		return
	}

	err := run()
	if err != nil {
		log.Error(err)
	}
}

func run() error {
	url := fmt.Sprintf("nats://%v:%v", hostname, port)
	if username != "" {
		url = fmt.Sprintf("nats://%v:%v@%v:%v", username, password, hostname, port)
	}

	nc, err := nats.Connect(url)
	if err != nil {
		return fmt.Errorf("failed to connect to NATS Server: %v", err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		return fmt.Errorf("failed to connect to JetStream: %v", err)
	}

	operation := flag.Arg(0)
	if operation == "" {
		return fmt.Errorf("missing operation, use either %s or %s", putFileHelp, getFileHelp)
	}

	objStoreName := flag.Arg(1)
	objStore, err := js.ObjectStore(objStoreName)
	if err != nil {
		objStore, err = js.CreateObjectStore(&nats.ObjectStoreConfig{
			Bucket: objStoreName,
		})
	}
	if err != nil {
		return fmt.Errorf("failed to bind/create object store: %v", err)
	}

	switch operation {
	case "put":
		_, err = objStore.PutFile(flag.Arg(2))
	case "get":
		err = objStore.GetFile(flag.Arg(2), flag.Arg(3))
	}

	return err
}
