// Package internal is an internal package
package internal

import (
	log "github.com/sirupsen/logrus"
)

// SetupLogging is used to setup logging
func SetupLogging(path string) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}
