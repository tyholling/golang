// Command server
package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/tyholling/golang/internal"
	"github.com/tyholling/golang/internal/grpc"
)

func main() {
	internal.SetupLogging()

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		server := http.Server{
			Addr:        ":8080",
			ReadTimeout: time.Minute,
		}
		log.Fatal(server.ListenAndServe())
	}()

	server := &grpc.Server{}
	err := server.Listen()
	if err != nil {
		log.Fatal(err)
	}
	server.Start()
}
