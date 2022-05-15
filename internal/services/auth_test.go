package services

import (
	"errors"
	"github.com/Zaida-3dO/goblin/config"
	"testing"

	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/internal/ports"
	"github.com/Zaida-3dO/goblin/internal/repositories"
	"github.com/Zaida-3dO/goblin/pkg/errs"
)

var as AuthService

func TestRegisterUserChecksForMissingFields(t *testing.T) {
	as = NewAuthService("test")

	mockErr := errs.NewBadRequestErr("some fields are missing", errors.New("confirm_password"))
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
		t.Errorf("expected refresh token to be empty string, got: %s\n", rt)
	}

	if err == nil {
		t.Errorf("expected err value, got nil\n")
	} else {
		if !err.Equals(mockErr) {
			t.Errorf("expected error=%+v, got=%+v\n", mockErr.ErrorDetails(), err.ErrorDetails())
		}
	}
}
func TestRegisterUserWithWrongEmailFormat(t *testing.T) {
	as = NewAuthService("test")

	mockErr := errs.NewBadRequestErr("some errors occurred", errors.New("invalid email address"))
	req := ports.RegisterUserRequest{
		FirstName:       "John",
		LastName:        "Doe",
		Email:           "johndoegmail.com",
		Password:        "johndoe",
		ConfirmPassword: "johndoe",
		PhoneNumber:     "09012345678",
	}

	usr, at, rt, err := as.RegisterUser(&req)

	if usr != nil {
		t.Errorf("expected user to be nil, got: %+v\n", usr)
	}
	if at != "" {
		t.Errorf("expected access token to be empty string, got: %s\n", at)
	}
	if rt != "" {
		t.Errorf("expected refresh token to be empty string, got: %s\n", rt)
	}

	if err == nil {
		t.Errorf("expected err value, got nil\n")
	} else {
		if !err.Equals(mockErr) {
			t.Errorf("expected error=%+v, got=%+v\n", mockErr.ErrorDetails(), err.ErrorDetails())
		}
	}
}
func TestRegisterUserWithPasswordNotEqualToConfirmPassword(t *testing.T) {
	as = NewAuthService("test")

	mockErr := errs.NewBadRequestErr("some errors occurred", errors.New("passwords do not match"))
	req := ports.RegisterUserRequest{
		FirstName:       "John",
		LastName:        "Doe",
		Email:           "johndoe@gmail.com",
		Password:        "johndoe",
		ConfirmPassword: "ohndoe",
		PhoneNumber:     "09012345678",
	}

	usr, at, rt, err := as.RegisterUser(&req)

	if usr != nil {
		t.Errorf("expected user to be nil, got: %+v\n", usr)
	}
	if at != "" {
		t.Errorf("expected access token to be empty string, got: %s\n", at)
	}
	if rt != "" {
		t.Errorf("expected refresh token to be empty string, got: %s\n", rt)
	}

	if err == nil {
		t.Errorf("expected err value, got nil\n")
	} else {
		if !err.Equals(mockErr) {
			t.Errorf("expected error=%+v, got=%+v\n", mockErr.ErrorDetails(), err.ErrorDetails())
		}
	}
}
func TestRegisterUserWithPasswordLengthLessThanSix(t *testing.T) {
	as = NewAuthService("test")

	mockErr := errs.NewBadRequestErr("some errors occurred", errors.New("password must be more than 6 characters"))
	req := ports.RegisterUserRequest{
		FirstName:       "John",
		LastName:        "Doe",
		Email:           "johndoe@gmail.com",
		Password:        "john",
		ConfirmPassword: "john",
		PhoneNumber:     "09012345678",
	}

	usr, at, rt, err := as.RegisterUser(&req)

	if usr != nil {
		t.Errorf("expected user to be nil, got: %+v\n", usr)
	}
	if at != "" {
		t.Errorf("expected access token to be empty string, got: %s\n", at)
	}
	if rt != "" {
		t.Errorf("expected refresh token to be empty string, got: %s\n", rt)
	}

	if err == nil {
		t.Errorf("expected err value, got nil\n")
	} else {
		if !err.Equals(mockErr) {
			t.Errorf("expected error=%+v, got=%+v\n", mockErr.ErrorDetails(), err.ErrorDetails())
		}
	}
}
func TestRegisterUserWithExistingEmail(t *testing.T) {
	as = NewAuthService("test")
	urm := repositories.UserRepoMock{}
	mockUser := dtos.NewUser("John", "Doe", "johndoe@gmail.com", "09012345678", "johndoe")

	mockErr := errs.NewBadRequestErr("email has been taken!", nil)
	req := ports.RegisterUserRequest{
		FirstName:       "John",
		LastName:        "Doe",
		Email:           "johndoe@gmail.com",
		Password:        "johndoe",
		ConfirmPassword: "johndoe",
		PhoneNumber:     "09012345678",
	}

	urm.CreateUser(*mockUser)
	defer urm.ResetDB()

	usr, at, rt, err := as.RegisterUser(&req)

	if usr != nil {
		t.Errorf("expected user to be nil, got: %+v\n", usr)
	}
	if at != "" {
		t.Errorf("expected access token to be empty string, got: %s\n", at)
	}
	if rt != "" {
		t.Errorf("expected refresh token to be empty string, got: %s\n", rt)
	}

	if err == nil {
		t.Errorf("expected err value, got nil\n")
	} else {
		if !err.Equals(mockErr) {
			t.Errorf("expected error=%+v, got=%+v\n", mockErr.ErrorDetails(), err.ErrorDetails())
		}
	}
}

func TestRegisterUserCreatesUser(t *testing.T) {
	config.LoadConfig("../../config")
	as = NewAuthService("test")

	mockUser := dtos.NewUser("John", "Doe", "johndoe@gmail.com", "09012345678", "johndoe")
	req := ports.RegisterUserRequest{
		FirstName:       "John",
		LastName:        "Doe",
		Email:           "johndoe@gmail.com",
		Password:        "johndoe",
		ConfirmPassword: "johndoe",
		PhoneNumber:     "09012345678",
	}

	usr, at, rt, err := as.RegisterUser(&req)

	if usr == nil {
		t.Errorf("expected user to be: %+v, got: nil\n", mockUser)
	}
	if at == "" {
		t.Errorf("expected access token to have value, got empty string\n")
	}
	if rt == "" {
		t.Errorf("expected refresh token to have value, got empty string\n")
	}

	if err != nil {
		t.Errorf("expected err value to be: nil got: %+v\n", err)
	} else {
		assertUserEqual(t, usr, mockUser)
	}
}

func assertUserEqual(t *testing.T, user *dtos.User, user2 *dtos.User) {
	if user.ID != user2.ID || user.Email != user2.Email || user.PhoneNumber != user2.PhoneNumber ||
		user.FirstName != user2.FirstName || user.LastName != user2.LastName {
		t.Errorf("expected=%+v, got=%+v\n", user2, user)
	}
}
