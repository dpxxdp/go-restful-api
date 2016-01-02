package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handling request from %v\n", r.RemoteAddr)
		fmt.Fprintf(w, "Hello, you've reached %q", html.EscapeString(r.URL.Path))
	})

	err := http.ListenAndServe(":8080", nil)

	if(err != nil) {
		log.Fatal("ListenAndServe: ", err)
	}
}