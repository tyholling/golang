// Command http
package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tyholling/golang/internal"
)

func main() {
	internal.SetupLogging()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	server := http.Server{
		ReadTimeout: time.Minute,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
