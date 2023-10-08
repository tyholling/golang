package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

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
