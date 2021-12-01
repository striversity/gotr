package main

import (
	"github.com/sirupsen/logrus"
	"github.com/striversity/misc007/moda/pkga"
	"github.com/striversity/misc007/moda/pkgc"
)

func main() {
	logrus.Infof("pi = %v\n", pkga.PI)
	logrus.Infof("secret key = %v\n", pkgc.SECRET_KEY)
}
