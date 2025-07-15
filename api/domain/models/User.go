package models

import (
	"errors"
	"net/mail"
	"unicode"
)

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Validated bool
}

func validatePassword(passowrd string) error {
	var hasNumber, hasSpecial bool

	for _, ch := range passowrd {
		switch {
		case unicode.IsNumber(ch):
			hasNumber = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	if !hasNumber && !hasSpecial {
		return errors.New("password needs numbers and special characters")
	}

	return nil
}

func (u User) Validate() error {
	if len(u.Username) < 5 {
		return errors.New("username must be at least 5 characters long.")
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return err
	}

	if len(u.Password) < 8 {
		return errors.New("password is too short.")
	}

	err = validatePassword(u.Password)
	if err != nil {
		return err
	}

	return nil
}
