package services

import (
	"time"

	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/internal/ports"
	"github.com/Zaida-3dO/goblin/internal/repositories"
	"github.com/Zaida-3dO/goblin/pkg/common"
	"github.com/Zaida-3dO/goblin/pkg/errs"
)

type AuthService interface {
	RegisterUser(req *ports.RegisterUserRequest) (*dtos.User, string, string, *errs.Err)
	LoginUser(req *ports.LoginRequest) (*dtos.User, string, string, *errs.Err)
}

type authService struct {
	userRepo      repositories.UserRepo
	userTokenRepo repositories.UserTokenRepo
}

func NewAuthService(mode string) AuthService {
	var as AuthService = &authService{
		userRepo:      repositories.NewUserRepo(mode),
		userTokenRepo: repositories.NewUserTokenRepo(mode),
	}
	return as
}

var ts = NewTokenService()

func (as *authService) RegisterUser(req *ports.RegisterUserRequest) (*dtos.User, string, string, *errs.Err) {
	err := req.ValidateRegisterUserRequest()
	if err != nil {
		return nil, "", "", err
	}

	var user = dtos.NewUser(req.FirstName, req.LastName, req.Email, req.PhoneNumber, req.Password)
	err = as.checkIfUserExists(req.Email)
	if err != nil {
		return nil, "", "", err
	}

	colour, colourErr := common.UserDefaultProfileColour(req.FirstName, req.LastName)
	if colourErr != nil {
		return nil, "", "", errs.NewBadRequestErr(colourErr.Error(), colourErr)
	}

	user.Colour = *colour
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	var hashErr error
	user.Password, hashErr = common.HashPassword(user.Password)
	if err != nil {
		return nil, "", "", errs.NewInternalServerErr(hashErr.Error(), hashErr)
	}

	err = as.userRepo.CreateUser(*user)
	if err != nil {
		return nil, "", "", err
	}

	var lr *ports.LoginRequest
	lr, err = ports.NewLoginRequest(req.Email, req.Password)
	if err != nil {
		return nil, "", "", err
	}

	return as.LoginUser(lr)
}

func (as *authService) LoginUser(req *ports.LoginRequest) (*dtos.User, string, string, *errs.Err) {
	var user dtos.User
	err := as.userRepo.FindUserByEmail(&user, req.Email)
	if err != nil {
		return nil, "", "", err
	}

	compareErr := common.ComparePassword(user.Password, req.Password)
	if compareErr != nil {
		return nil, "", "", errs.NewBadRequestErr("invalid login credentials", compareErr)
	}

	var tokenPair *token
	tokenPair, err = ts.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, "", "", err
	}

	userToken := dtos.NewUserToken(user.ID, tokenPair.accessToken, tokenPair.refreshToken, tokenPair.accessUUID, tokenPair.refreshUUID)

	err = as.userTokenRepo.CreateToken(*userToken)
	if err != nil {
		return nil, "", "", err
	}

	return &user, tokenPair.accessToken, tokenPair.refreshToken, nil
}

func (as *authService) checkIfUserExists(email string) *errs.Err {
	var user dtos.User
	err := as.userRepo.FindUserByEmail(&user, email)
	if err == nil {
		return errs.NewBadRequestErr("email has been taken!", nil)
	}
	return nil
}
