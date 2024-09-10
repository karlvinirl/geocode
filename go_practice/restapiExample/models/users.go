package models

import (
	"errors"

	"example.com/go_practice/restapi/db"
	"example.com/go_practice/restapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string

	//we bind the user ID here from DB to ensure user had valid id post validation.
	err := row.Scan(&user.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	isValid := utils.CheckPasswordHash(user.Password, retrievedPassword)
	if !isValid {
		return errors.New("credentials invalid")
	}

	return nil
}
func (user *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	user.ID = userId
	return err
}
