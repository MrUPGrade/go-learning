package main

import (
	"fmt"
	"log"
	"net/http"
)

type api struct{}

func (api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world from %v", r.URL.Path)
	if err != nil {
		log.Print(err)
	}
}

func main() {
	a := api{}
	s := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: a,
	}

	err := s.ListenAndServe()
	log.Fatal(err)
}
