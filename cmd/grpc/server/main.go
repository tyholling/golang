// Command server
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/tyholling/golang/internal"
	"github.com/tyholling/golang/internal/grpc"
)

func main() {
	internal.SetupLogging()

	server := &grpc.Server{}
	err := server.Listen()
	if err != nil {
		log.Fatal(err)
	}
	server.Start()
}
