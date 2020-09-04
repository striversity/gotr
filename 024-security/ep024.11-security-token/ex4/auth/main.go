package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var (
	userDb *UserDB
)

func main() {
	app := newApp()
	app.Action = startWebServer
	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "addr", Value: ":8081", Usage: "server socket bind interface and port"},
	}

	userDb = &UserDB{}
	userDb.init()

	err := app.Run(os.Args)
	if err != nil {
		logrus.Error(err)
	}
}

func startWebServer(ctx *cli.Context) error {
	addr := ctx.String("addr")
	fmt.Printf("Listening on %v\n", addr)

	http.HandleFunc("/login", loginHandler)

	return http.ListenAndServe(addr, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("Hello, login requested from %v", r.RemoteAddr)
}

func newApp() *cli.App {
	app := &cli.App{}
	app.Name = "auth server"
	app.Description = "auth server is my simple authentication/authorization webserver application"
	app.Version = "1"
	app.Authors = []*cli.Author{
		{Name: "Striversity, LLC", Email: "youtube@striversity.com"},
	}

	return app
}
