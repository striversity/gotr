package main

import (
  "flag"
  "fmt"
  "log"
  "time"

  "github.com/nats-io/nats.go"
)

var (
  username string
  password string
  hostname = "localhost"
  port     = 4222
)

func init() {
  flag.StringVar(&username, "u", username, "username for NATS Server")
  flag.StringVar(&password, "p", password, "password for NATS Server")
  flag.StringVar(&hostname, "host", hostname, "NAST Server hostname")
  flag.IntVar(&port, "port", port, "NAST Server port")
  flag.Parse()
}

func fatalOnErr(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func main() {
  url := fmt.Sprintf("nats://%v:%v", hostname, port)
  if username != "" {
    url = fmt.Sprintf("nats://%v:%v@%v:%v", username, password, hostname, port)
  }

  nc, err := nats.Connect(url)
  fatalOnErr(err)
  defer nc.Close()

  js, err := nc.JetStream()
  fatalOnErr(err)

  _, err = js.AddConsumer("ORDERS", &nats.ConsumerConfig{
    Durable:      "my-consumer-1",
    Description:  "this is my awesome consumer",
    ReplayPolicy: nats.ReplayInstantPolicy,
  })
  fatalOnErr(err)

  sub, err := js.PullSubscribe("orders.us", "my-consumer-1")
  fatalOnErr(err)
  go processMsg(sub)

  time.Sleep(10 * time.Second)
  sub.Unsubscribe()

  log.Println("shutting down application...")
}

func processMsg(sub *nats.Subscription) {
  for sub.IsValid() {
    msgs, err := sub.Fetch(1)
    if err == nil {
      fmt.Printf("INFO - Got reply - %s\n", string(msgs[0].Data))
    }
  }
}

