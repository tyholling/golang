// Package internal is an internal package
package internal

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// SetupLogging is used to setup logging
func SetupLogging(path string) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	err := os.MkdirAll("log", 0o0644)
	if err != nil {
		log.Fatalf("failed to create log directory: %s", err)
	}
	file, err := os.Create(filepath.Clean("log/" + path))
	if err != nil {
		log.Fatalf("failed to create log file: %s", err)
	} else {
		log.SetOutput(file)
	}
}
