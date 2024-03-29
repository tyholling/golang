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

	http.HandleFunc("/", handle)

	server := http.Server{
		ReadTimeout: time.Minute,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

	data := []byte("\n")
	_, err := w.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
