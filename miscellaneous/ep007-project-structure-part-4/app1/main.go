package main

import (
	"moda/pkga"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("pi = %v\n", pkga.PI)
}
