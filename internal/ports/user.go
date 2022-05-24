package ports

import (
	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/pkg/common"
	"github.com/Zaida-3dO/goblin/pkg/errs"
)

type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (cpr *ChangePasswordRequest) ValidateChangePasswordRequest() *errs.Err {
	keyBindings := []string{"old_password", "new_password", "confirm_password"}

	err := common.ValidateHttpRequestsForMissingFields(cpr, ChangePasswordRequest{}, keyBindings)
	if err != nil {
		return err
	}

	_, pwdErr := common.NewPassword(cpr.NewPassword, cpr.ConfirmPassword)
	if pwdErr != nil {
		return errs.NewBadRequestErr(pwdErr.Error(), pwdErr)
	}

	return nil
}

func ChangePasswordReply() *map[string]interface{} {
	return &map[string]interface{}{
		"message": "password changed successfully",
	}
}

type UpdateUserRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
}

func (uur *UpdateUserRequest) ValidateUpdateUserRequest() *errs.Err {
	keyBindings := []string{"first_name", "last_name", "email", "phone_number", "gender"}

	err := common.ValidateHttpRequestsForMissingFields(uur, UpdateUserRequest{}, keyBindings)
	if err != nil {
		return err
	}

	email, emailErr := common.NewEmail(uur.Email)
	if emailErr != nil {
		return errs.NewBadRequestErr(emailErr.Error(), emailErr)
	}

	uur.Email = email.Address

	return nil
}

func UpdateUserReply(user dtos.User) *map[string]interface{} {
	return &map[string]interface{}{
		"message": "your details have been successfully updated",
		"user": user.Strip(),
	}
}
