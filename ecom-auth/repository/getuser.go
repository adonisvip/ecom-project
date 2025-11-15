package repository

import (
	"database/sql"
	db "ecom-auth/pkg/postgres"
	"ecom-auth/pkg/utils"
	"errors"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Fullname sql.NullString
}

func GetUserByUsername(username string) (*User, error) {
	query := "SELECT id, username, password, email, fullname FROM users WHERE username = $1"
	row := db.DB.QueryRow(query, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Fullname)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func UpdateUserTokens(userID int, token, refreshToken string) error {
	query := "UPDATE users SET token = $1, refresh_token = $2 WHERE id = $3"
	_, err := db.DB.Exec(query, token, refreshToken, userID)
	return err
}

func GetUserByRefreshToken(refreshToken string) (*User, error) {
	query := "SELECT id, username, password, email, fullname FROM users WHERE refresh_token = $1"
	row := db.DB.QueryRow(query, refreshToken)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Fullname)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("refresh token not found")
		}
		return nil, err
	}
	return &user, nil
}

func UpdateAccessTokenOnly(userID int, token string) error {
	query := "UPDATE users SET token = $1 WHERE id = $2"
	_, err := db.DB.Exec(query, token, userID)
	return err
}

func GetUserByToken(token string) (*User, error) {
	query := "SELECT id, username, password, email, fullname FROM users WHERE token = $1"
	row := db.DB.QueryRow(query, token)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Fullname)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("token not found")
		}
		return nil, err
	}
	return &user, nil
}

func UserExists(username string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE username = $1"
	var count int
	err := db.DB.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func EmailExists(email string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	var count int
	err := db.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CreateUser(username, password, email string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3)"
	_, err = db.DB.Exec(query, username, hashedPassword, email)
	return err
}
