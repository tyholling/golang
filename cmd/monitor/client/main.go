// command client
package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout
}

func main() {
	log.Info("client: started")

	log.Info("client: stopped")
}
