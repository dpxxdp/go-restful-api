package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"database/sql"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

type Env struct {
	db *sql.DB
}

func main() {
	pass := os.Getenv("PSPASS")

	db, dbErr := sql.Open("postgres", fmt.Sprintf("user=postgres password=%v dbname=gorestapp sslmode=disable", pass))
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	env := &Env{db: db}

	router := httprouter.New()

	router.GET("/", IndexGet)

	router.GET("/users", UsersGet(env))
	router.GET("/medias", MediasGet)
	router.GET("/channels", ChannelsGet)

	router.GET("/users/:id", UserGet)
	router.GET("/medias/:id", MediaGet)
	router.GET("/channels/:id", ChannelGet)

	loggingRouter := Logger(router)

	err := http.ListenAndServe(":8080", loggingRouter)

	if(err != nil) {
		log.Fatal("ListenAndServe: ", err)
	}
}
