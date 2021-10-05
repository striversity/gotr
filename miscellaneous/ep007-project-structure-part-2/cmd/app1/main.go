package main

import (
	"github.com/sirupsen/logrus"
	"github.com/striversity/gotr/misc-ep007/feata"
	"github.com/striversity/gotr/misc-ep007/featb"
)

func main() {
	if err := feata.Do(); err != nil {
		logrus.Fatal(err)
	}

	if err := featb.Do(); err != nil {
		logrus.Fatal(err)
	}
}
