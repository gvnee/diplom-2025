package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("GET	/{$}", home)
	mux.HandleFunc("GET	/snippet/view/{id}", snippetView)
	mux.HandleFunc("GET	/snippet/create", snippetCreate)
	mux.HandleFunc("POST /test", test)

	log.Printf("starting server on %s", *addr)

	handler := cors.Default().Handler(mux)

	err := http.ListenAndServe(*addr, handler)
	log.Fatal(err)
}
