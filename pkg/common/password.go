package common

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Value string
}

func NewPassword(original, confirmed string) (*Password, error) {
	length := len([]rune(original))
	if length < 6 {
		return nil, errors.New("password must be more than 6 characters")
	}
	if original != confirmed {
		return nil, errors.New("passwords do not match")
	}
	return &Password{
		Value: original,
	}, nil
}

func HashPassword(passwordValue string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(passwordValue), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error processing string")
	}
	return string(passwordHash), nil
}

func ComparePassword(hashedPassword string, password string) error {
	pw := []byte(password)
	hpw := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(hpw, pw)
	return err
}
