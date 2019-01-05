package main

import (
	"log"
	"net/http"
)

func main() {
	root := func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hello world"))
		if err != nil {
			log.Print(err)
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", root)

	s := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := s.ListenAndServe()
	log.Fatal(err)
}
