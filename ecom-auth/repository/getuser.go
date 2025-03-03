package repository

import (
	"database/sql"
	"errors"
	db "ecom-auth/pkg/postgres"
	"ecom-auth/pkg/utils"
)

type User struct {
	ID       int
	Username string
	Password string
}

func GetUserByUsername(username string) (*User, error) {
	query := "SELECT id, username, password FROM users WHERE username = $1"
	row := db.DB.QueryRow(query, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(username, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err = db.DB.Exec(query, username, hashedPassword)
	return err
}