package ports

import (
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
