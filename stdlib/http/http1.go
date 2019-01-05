package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello world from %v", r.URL.Path)
		if err != nil {
			log.Print(err)
		}
	})

	s := http.Server{
		Addr: "0.0.0.0:8080",
	}

	err := s.ListenAndServe()
	log.Fatal(err)
}
