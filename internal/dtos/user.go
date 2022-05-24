package dtos

import (
	"time"
)

type APIUser struct {
	ID                           uint      `json:"id"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	FirstName                    string    `json:"first_name"`
	LastName                     string    `json:"last_name"`
	Email                        string    `json:"email"`
	PhoneNumber                  string    `json:"phone_number"`
	Colour                       string    `json:"colour"`
	Gender                       string    `json:"gender"`
}

func NewAPIUser(
	id uint,
	createdAt,
	updatedAt time.Time,
	firstName,
	lastName,
	email,
	phoneNumber,
	colour,
	gender string,
) APIUser {
	return APIUser{
		ID: id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		PhoneNumber:     phoneNumber,
		Colour: colour,
		Gender: gender,
	}
}

type User struct {
	ID                           uint      `json:"id"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	FirstName                    string    `json:"first_name"`
	LastName                     string    `json:"last_name"`
	Email                        string    `json:"email"`
	PhoneNumber                  string    `json:"phone_number"`
	Password                     string    `json:"password"`
	EmailVerificationCode        string    `json:"email_verification_code"`
	IsEmailVerified              bool      `json:"is_email_verified"`
	PasswordResetToken           string    `json:"password_reset_token"`
	PasswordResetTokenExpiryDate string    `json:"password_reset_token_expiry_date"`
	Colour                       string    `json:"colour"`
	Gender                       string    `json:"gender"`
}

func NewUser(
	firstName,
	lastName,
	email,
	phoneNumber,
	password string,
) *User {
	return &User{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		PhoneNumber:     phoneNumber,
		Password:        password,
		IsEmailVerified: false,
	}
}

func (u *User) Strip() APIUser {
	return NewAPIUser(u.ID, u.CreatedAt, u.UpdatedAt, u.FirstName,
		u.LastName, u.Email, u.PhoneNumber, u.Colour, u.Gender)
}
