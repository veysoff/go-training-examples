package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handler_1")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Handler_2")
}

func main() {
	http.HandleFunc("/endpoint1", logging(handler1))
	http.HandleFunc("/endpoint2", logging(handler2))

	http.ListenAndServe(":8080", nil)
}
