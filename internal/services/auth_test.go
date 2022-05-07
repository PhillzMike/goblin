package services

import (
	"testing"

	"github.com/Zaida-3dO/goblin/internal/ports"
	"github.com/Zaida-3dO/goblin/pkg/errs"
)

var as AuthService

func TestRegisterUserChecksForMissingFields(t *testing.T) {
	as = NewAuthService("test")

	mockErr := errs.NewBadRequestErr("some fields are missing", []string{"confirm_password"})
	req := ports.RegisterUserRequest{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "johndoe@gmail.com",
		Password:    "johndoe",
		PhoneNumber: "09012345678",
	}

	usr, at, rt, err := as.RegisterUser(&req)

	if usr != nil {
		t.Errorf("expected user to be nil, got: %+v\n", usr)
	}
	if at != "" {
		t.Errorf("expected access token to be empty string, got: %s\n", at)
	}
	if rt != "" {
		t.Errorf("expected refresh token to be empty string, got: %s\n", at)
	}

	if err == nil {
		t.Errorf("expected access token to be empty string, got: %s\n", at)
	} else {
		if !err.Equals(mockErr) {
			t.Errorf("expected error=%+v, got=%+v\n", mockErr.ErrorDetails(), err.ErrorDetails())
		}
	}
}
func TestRegisterUserWithWrongEmailFormat(t *testing.T)                  {}
func TestRegisterUserWithPasswordNotEqualToConfirmPassword(t *testing.T) {}
func TestRegisterUserWithPasswordLengthLessThanSix(t *testing.T)         {}
func TestRegisterUserWithExistingEmail(t *testing.T)                     {}
func TestRegisterUserCreatesUser(t *testing.T)                           {}
