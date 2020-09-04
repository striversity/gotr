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

	userDb = &UserDB{}
	userDb.init()

	err := app.Run(os.Args)
	if err != nil {
		logrus.Error(err)
	}
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
