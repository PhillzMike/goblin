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
	err := errs.NewBadRequestErr("some errors occurred", nil)
	keyBindings := []string{"first_name", "last_name", "email", "password", "confirm_password"}

	reqErr := common.ValidateHttpRequestsForMissingFields(rur, RegisterUserRequest{}, keyBindings)
	if reqErr != nil {
		return reqErr
	}

	em, newEmailErr := common.NewEmail(rur.Email)
	if newEmailErr != nil {
		err.Add(newEmailErr)
	}

	pwd, newPasswordErr := common.NewPassword(rur.Password, rur.ConfirmPassword)
	if newPasswordErr != nil {
		err.Add(newPasswordErr)
	}

	if err.HasData() {
		return err
	}

	rur.Email = em.Address
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

func LoginReply(user *dtos.User, accessToken, refreshToken, message string) *map[string]interface{} {
	return &map[string]interface{}{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"message":       message,
	}
}

type ForgotPasswordRequest struct {
	Email      string `json:"email"`
	RedirectTo string `json:"redirect_to"`
}

func (fpr *ForgotPasswordRequest) ValidateForgotPasswordRequest() *errs.Err {
	keyBindings := []string{"email", "redirect_to"}

	err := common.ValidateHttpRequestsForMissingFields(fpr, ForgotPasswordRequest{}, keyBindings)
	if err != nil {
		return err
	}

	emailObject, newEmailErr := common.NewEmail(fpr.Email)
	if newEmailErr != nil {
		return errs.NewBadRequestErr(newEmailErr.Error(), newEmailErr)
	}
	fpr.Email = emailObject.Address

	return nil
}

func ForgotPasswordReply() *map[string]string {
	return &map[string]string{
		"message": "check your email for instructions to retrieve your password",
	}
}

type ResetPasswordRequest struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Token           string `json:"token"`
}

func (rpr *ResetPasswordRequest) ValidateResetPasswordRequest() *errs.Err {
	keyBindings := []string{"password", "confirm_password", "token"}

	err := common.ValidateHttpRequestsForMissingFields(rpr, ResetPasswordRequest{}, keyBindings)
	if err != nil {
		return err
	}

	pwd, pwdErr := common.NewPassword(rpr.Password, rpr.ConfirmPassword)
	if pwdErr != nil {
		return err
	}
	rpr.Password = pwd.Value

	return nil
}

func ResetPasswordReply() *map[string]string {
	return &map[string]string{
		"message": "your password has been reset successfully",
	}
}
