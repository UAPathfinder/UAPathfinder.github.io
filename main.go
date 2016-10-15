package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
)

var (
	listen = flag.String("listen", ":8080", "The adress this service will be available on.")
)

func main() {
	flag.Parse()
	mux := http.NewServeMux()

	staticFiles := rice.MustFindBox("frontend").HTTPBox()
	mux.Handle("/", http.FileServer(staticFiles))

	mux.HandleFunc("/api/courses", func(rw http.ResponseWriter, r *http.Request) {
		// TODO: Implement Course Serialization
	})

	mux.HandleFunc("/api/combos", func(rw http.ResponseWriter, r *http.Request) {
		// TODO: Get a number of combos along with their scores.
	})

	server := &http.Server{
		Addr:    *listen,
		Handler: mux,
	}

	log.Printf("Starting server on %s\n", *listen)
	log.Fatalln(server.ListenAndServe())
}
