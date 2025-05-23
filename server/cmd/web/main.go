package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("GET	/{$}", home)
	mux.HandleFunc("GET	/snippet/view/{id}", snippetView)
	mux.HandleFunc("GET	/snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("starting server on %s", *addr)

	runC()

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
