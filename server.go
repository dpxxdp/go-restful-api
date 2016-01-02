package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello and welcome!")
}

func User(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "user %v", params[0])
}

func Text(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "text %v", params[0])
}

func Channel(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "channel %v", params[0])
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/user/:id", User)
	router.GET("/text/:id", Text)
	router.GET("/channel/:id", Channel)

	err := http.ListenAndServe(":8080", router)

	if(err != nil) {
		log.Fatal("ListenAndServe: ", err)
	}
}