package main

import (
	"time"
)

type User struct {
	Id			int			`json:"id"`
	Name 		string		`json:"name"`
	Email 		string		`json:"email"`
	Created 	time.Time 	`json:"created"`
	Hash		string		`json:"-"`
	Salt		string		`json:"-"`
	Active		bool		`json:"active"`
}

type Users []User

type Channel struct {
	Id			int			`json:"id"`
	Name		string		`json:"name"`
	Created 	time.Time 	`json:"created"`
	Mods		Users		`json:"mods"`
	Users		Users		`json:"users"`
	Active 		bool 		`json:"active"`
}

type Channels []Channel

type Media struct {
	Id			int			`json:"id"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	User		User 		`json:"user"`
	Channel		Channel 	`json:"channel"`
	Created		time.Time 	`json:"time"`
	Points		int			`json:"points"`
	Active		bool 		`json:"active"`
}

type Medias []Media