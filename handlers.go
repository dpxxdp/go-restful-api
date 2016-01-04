package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func IndexGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello and welcome!\n")
}

func UsersGet(env *Env) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		users, dbErr := DalGetAllUsers(env.db)
		var emptyUser struct{}

		if dbErr != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(emptyUser)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
	
			if err := json.NewEncoder(w).Encode(users); err != nil {
    	    	panic(err)
    		}
    	}
    })
}

func UserGet(env *Env) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var emptyUser struct{}

		id, err := strconv.Atoi(params[0].Value)
		if(err != nil) {
			fmt.Printf("%v", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(emptyUser)
		}

		user, dbErr := DalGetUser(id, env.db)

		if dbErr != nil {
			fmt.Printf("%v", dbErr)
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(emptyUser)
		} else {
			w.WriteHeader(http.StatusOK)
	
			if err := json.NewEncoder(w).Encode(user); err != nil {
    	    	panic(err)
    		}
    	}
    })
}

func UserPost(env *Env) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var emptyUser struct{}

		decoder := json.NewDecoder(r.Body)
    	var user User
    	err := decoder.Decode(&user)
    	if err != nil {
        	w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(emptyUser)
    	}

		returnUser, dbErr := DalCreateUser(&user, env.db)

		if dbErr != nil {
			fmt.Printf("%v", dbErr)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(emptyUser)
		} else {
			w.WriteHeader(http.StatusCreated)
	
			if err := json.NewEncoder(w).Encode(returnUser); err != nil {
    	    	panic(err)
    		}
    	}
    })
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