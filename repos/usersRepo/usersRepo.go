package usersRepo

import (
	"database/sql"
	"fmt"
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

func GetUserByUname(username string) (*ORMs.User, error) {
	query := "SELECT userId, username, password FROM users WHERE username = @username"

	var user ORMs.User
	err := data.DB.QueryRow(query, sql.Named("username", username)).Scan(&user.UserId, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("User not found")
			return nil, nil // User not found
		}
		fmt.Println("Error querying user:", err)
		return nil, err // Other errors
	}
	return &user, nil
}

