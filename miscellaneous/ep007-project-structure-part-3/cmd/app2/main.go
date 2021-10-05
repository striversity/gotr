package main

import (
	"github.com/sirupsen/logrus"
	"github.com/striversity/gotr/misc-ep007/internal/cmd/app2"
)

func main() {
	if err := app2.Run(); err != nil {
		logrus.Fatal(err)
	}
}
