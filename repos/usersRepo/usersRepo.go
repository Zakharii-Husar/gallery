package usersRepo

import (
	"database/sql"
	"gallery/data"
	"gallery/models/ORMs"
	"log"
)

func CreateUser(input ORMs.User) error {
    query := "INSERT INTO users (username, password) VALUES (@Username, @Password)"
    _, err := data.DB.Exec(query, sql.Named("Username", input.Username), sql.Named("Password", input.Password))
    if err != nil {
        log.Println("Error inserting user: ", err)
        return err
    }
    return nil
}
