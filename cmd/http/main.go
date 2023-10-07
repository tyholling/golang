package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handle)

	server := http.Server{
		ReadTimeout: time.Minute,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func handle(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

	data := []byte("\n")
	_, err := w.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
}
