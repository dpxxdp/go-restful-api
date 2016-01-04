package main

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
)

func DalGetAllUsers(db *sql.DB) ([]*User, error) {

	rows, err := db.Query("SELECT user_id, name, email, created, active FROM users")
    if err != nil {
    	fmt.Printf("%v", err)
        return nil, err
    }

    defer rows.Close()

    users := make([]*User, 0)

    for rows.Next() {
        user := new(User)
        err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Created, &user.Active)
        if err != nil {
        	fmt.Printf("%v", err)
            return nil, err
        }
        users = append(users, user)
    }
    if err = rows.Err(); err != nil {
    	fmt.Printf("%v", err)
        return nil, err
    }
    return users, nil
}

func DalGetUser(id int, db *sql.DB) (*User, error) {

	rows, err := db.Query("SELECT user_id, name, email, created, active FROM users where user_id = $1", id)
    if err != nil {
    	fmt.Printf("%v", err)
        return nil, err
    }

    defer rows.Close()

    user := new(User)

    rows.Next()
    scanErr := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Created, &user.Active)
    if scanErr != nil {
    	fmt.Printf("%v", scanErr)
        return nil, scanErr
    }
    if err = rows.Err(); err != nil {
    	fmt.Printf("%v", err)
        return nil, err
    }
    return user, nil
}

func DalCreateUser(user *User, db *sql.DB) (*User, error) {

	_, err := db.Exec("INSERT INTO users(user_id, name, email, created, active) VALUES($1, $2, $3, $4, $5)", user.Id, user.Name, user.Email, user.Created, user.Active)
	if err != nil {
		return nil, err
	}

    return user, nil
}

