// Command client
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/tyholling/golang/internal"
	"github.com/tyholling/golang/internal/grpc"
)

func main() {
	internal.SetupLogging()

	client := &grpc.Client{}
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	client.Start()
}
