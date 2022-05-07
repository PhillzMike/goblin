package ports

import (
	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/pkg/common"
	"github.com/Zaida-3dO/goblin/pkg/errs"
)

type RegisterUserRequest struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	PhoneNumber     string `json:"phone_number"`
}

func (rur *RegisterUserRequest) ValidateRegisterUserRequest() *errs.Err {
	keyBindings := []string{"first_name", "last_name", "email", "password", "confirm_password"}

	err := common.ValidateHttpRequestsForMissingFields(rur, RegisterUserRequest{}, keyBindings)
	if err != nil {
		return err
	}

	em, newEmailErr := common.NewEmail(rur.Email)
	if newEmailErr != nil {
		return errs.NewBadRequestErr(newEmailErr.Error(), newEmailErr)
	}
	rur.Email = em.Address

	pwd, newPasswordErr := common.NewPassword(rur.Password, rur.ConfirmPassword)
	if newPasswordErr != nil {
		return errs.NewBadRequestErr(newPasswordErr.Error(), newPasswordErr)
	}
	rur.Password = pwd.Value

	return nil
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLoginRequest(email, password string) (*LoginRequest, *errs.Err) {
	lr := &LoginRequest{
		Email:    email,
		Password: password,
	}
	err := lr.ValidateLoginRequest()
	if err != nil {
		return nil, err
	}
	return lr, nil
}

func (lr *LoginRequest) ValidateLoginRequest() *errs.Err {
	keyBindings := []string{"email", "password"}

	err := common.ValidateHttpRequestsForMissingFields(lr, LoginRequest{}, keyBindings)
	if err != nil {
		return err
	}

	emailObject, newEmailErr := common.NewEmail(lr.Email)
	if newEmailErr != nil {
		return errs.NewBadRequestErr(newEmailErr.Error(), newEmailErr)
	}
	lr.Email = emailObject.Address

	return nil
}

type LoginReply map[string]interface{}

func NewLoginReply(user *dtos.User, accessToken, refreshToken, message string) *LoginReply {
	return &LoginReply{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"message":       message,
	}
}
