package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

func IndexGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello and welcome!\n")
}

func UsersGet(env *Env) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		users, dbErr := AllUsers(env.db)

		if dbErr != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "<h1>404</h1><div>Not Found</div>")
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
	
			if err := json.NewEncoder(w).Encode(users); err != nil {
    	    	panic(err)
    		}
    	}
    })
}

func UserGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	
	users := Users{
		User{Id: 0, Name: "dan"},
		User{Id: 1, Name: "ada"},
		User{Id: 2, Name: "bob"},
		User{Id: 3, Name: "adam"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
       	panic(err)
    }
}

func UserPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func MediasGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "media %v", params[0])
}

func MediaGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "media %v", params[0])
}

func ChannelsGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "channel %v", params[0])
}

func ChannelGet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "channel %v", params[0])
}