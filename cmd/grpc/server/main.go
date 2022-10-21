// Command server
package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/tyholling/golang/internal/grpc"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	err := os.MkdirAll("log", 0o0644)
	if err != nil {
		log.Fatalf("failed to create log directory: %s", err)
	}
	file, err := os.Create("log/server.log")
	if err != nil {
		log.Fatalf("failed to create log file: %s", err)
	} else {
		log.SetOutput(file)
	}
}

func main() {
	server := &grpc.Server{}
	err := server.Listen()
	if err != nil {
		log.Fatal(err)
	}
	server.Start()
}
