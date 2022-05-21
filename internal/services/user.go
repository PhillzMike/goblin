package services

import (
	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/internal/ports"
	"github.com/Zaida-3dO/goblin/internal/repositories"
	"github.com/Zaida-3dO/goblin/pkg/common"
	"github.com/Zaida-3dO/goblin/pkg/errs"
)

type UserService interface {
	ChangePassword(*ports.ChangePasswordRequest, *dtos.User) *errs.Err
}

type userService struct {
	userRepo     repositories.UserRepo
	emailService EmailServiceInterface
}

func NewUserService(mode string) UserService {
	var us UserService = &userService{
		userRepo:     repositories.NewUserRepo(mode),
		emailService: NewEmailService(),
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

	return nil
}
