// Package internal is an internal package
package internal

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// SetupLogging is used to setup logging.
func SetupLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}
