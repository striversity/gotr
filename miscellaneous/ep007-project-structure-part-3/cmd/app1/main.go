package main

import (
	"github.com/sirupsen/logrus"
	"github.com/striversity/gotr/misc-ep007/internal/cmd/app1"
)

func main() {
	if err := app1.Run(); err != nil {
		logrus.Fatal(err)
	}
}
