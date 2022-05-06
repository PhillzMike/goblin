package common

import (
	"errors"
	"net/mail"
)

type Email struct {
	Address string
}

func NewEmail(address string) (*Email, error) {
	addr, b := validateEmailSyntax(address)
	if !b {
		return nil, errors.New("invalid email address")
	}
	return &Email{
		Address: *addr,
	}, nil
}

func validateEmailSyntax(address string) (*string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return nil, false
	}
	return &addr.Address, true
}
