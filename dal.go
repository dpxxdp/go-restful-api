package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func AllUsers(db *sql.DB) ([]*User, error) {

	rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    users := make([]*User, 0)

    for rows.Next() {
        user := new(User)
        err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Created, &user.Hash, &user.Salt, &user.Active)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    if err = rows.Err(); err != nil {
        return nil, err
    }
    return users, nil
}

