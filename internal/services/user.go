package services

import (
	"fmt"

	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/internal/ports"
	"github.com/Zaida-3dO/goblin/internal/repositories"
	"github.com/Zaida-3dO/goblin/pkg/common"
	"github.com/Zaida-3dO/goblin/pkg/errs"
)

type UserService interface {
	ChangePassword(*ports.ChangePasswordRequest, *dtos.User) *errs.Err
	UpdateUser(*ports.UpdateUserRequest, *dtos.User) *errs.Err
	DeleteUser(*dtos.User) *errs.Err
}

type userService struct {
	userRepo repositories.UserRepo
}

func NewUserService(mode string) UserService {
	var us UserService = &userService{
		userRepo: repositories.NewUserRepo(mode),
	}
	return us
}

func (us *userService) ChangePassword(req *ports.ChangePasswordRequest, currentUser *dtos.User) *errs.Err {
	err := req.ValidateChangePasswordRequest()
	if err != nil {
		return err
	}

	compareErr := common.ComparePassword(currentUser.Password, req.OldPassword)
	if compareErr != nil {
		return errs.NewBadRequestErr("old password is not correct", nil)
	}

	var hashErr error
	currentUser.Password, hashErr = common.HashPassword(req.NewPassword)
	if err != nil {
		return errs.NewInternalServerErr(hashErr.Error(), hashErr)
	}

	us.userRepo.SaveUser(currentUser)

	es := NewEmailService()

	if err := es.SendChangedPasswordEmail(currentUser.FirstName, currentUser.Email); err != nil {
		// log the error
		fmt.Printf("error sending email: %v\n", err)
	}

	return nil
}

func (us *userService) UpdateUser(req *ports.UpdateUserRequest, currentUser *dtos.User) *errs.Err {
	err := req.ValidateUpdateUserRequest()
	if err != nil {
		return err
	}

	_, err = EnsureUserIsNotCurrentUserIfExists(us.userRepo, currentUser, req.Email)
	if err != nil {
		return err
	}

	currentUser.FirstName = req.FirstName
	currentUser.LastName = req.LastName
	currentUser.Email = req.Email
	currentUser.PhoneNumber = req.PhoneNumber
	currentUser.Gender = req.Gender

	if err = us.userRepo.SaveUser(currentUser); err != nil {
		return err
	}

	return nil
}

func (us *userService) DeleteUser(user *dtos.User) *errs.Err {
	if err := us.userRepo.DeleteUser(*user); err != nil {
		return err
	}
	return nil
}

func EnsureEmailNotTaken(repo repositories.UserRepo, email string) *errs.Err {
	var user dtos.User
	if err := repo.FindUserByEmail(&user, email); err == nil {
		return errs.NewBadRequestErr("email has been taken!", nil)
	}
	return nil
}

func EnsureUserIsNotCurrentUserIfExists(repo repositories.UserRepo, currentUser *dtos.User, email string) (*dtos.User, *errs.Err) {
	var user dtos.User
	if err := repo.FindUserByEmail(&user, email); err == nil && user.ID != currentUser.ID {
		return nil, errs.NewBadRequestErr("email has been taken!", nil)
	}
	return &user, nil
}
