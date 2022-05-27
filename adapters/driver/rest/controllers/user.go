package controllers

import (
	"errors"
	"net/http"

	"github.com/Zaida-3dO/goblin/internal/dtos"
	"github.com/Zaida-3dO/goblin/internal/ports"
	"github.com/Zaida-3dO/goblin/internal/services"
	"github.com/Zaida-3dO/goblin/pkg/errs"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	ChangePassword(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	userService services.UserService
	authService services.AuthService
}

func NewUserController(mode string) UserController {
	var uc UserController = &userController{
		userService: services.NewUserService(mode),
		authService: services.NewAuthService(mode),
	}
	return uc
}

func (uc *userController) ChangePassword(c *gin.Context) {
	user, b := GetCurrentUser(c)
	if !b {
		return
	}

	var request ports.ChangePasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		httpErr := errs.NewBadRequestErr("invalid json body", err)
		c.JSON(httpErr.StatusCode, httpErr)
		return
	}

	if err := uc.userService.ChangePassword(&request, user); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	response := ports.ChangePasswordReply()

	c.JSON(http.StatusOK, response)
}

func (uc *userController) UpdateUser(c *gin.Context) {
	user, b := GetCurrentUser(c)
	if !b {
		return
	}

	var request ports.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		httpErr := errs.NewBadRequestErr("invalid json body", err)
		c.JSON(httpErr.StatusCode, httpErr)
		return
	}

	if err := uc.userService.UpdateUser(&request, user); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	response := ports.UpdateUserReply(*user)

	c.JSON(http.StatusOK, response)
}

func (uc *userController) DeleteUser(c *gin.Context) {
	user, b := GetCurrentUser(c)
	if !b {
		return
	}

	var request ports.DeleteUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		httpErr := errs.NewBadRequestErr("invalid json body", err)
		c.JSON(httpErr.StatusCode, httpErr)
		return
	}

	var lr ports.LoginRequest
	lr.Email = user.Email
	lr.Password = request.Password

	_, _, _, err := uc.authService.LoginUser(&lr)
	if err != nil {
		loginErr := errs.NewBadRequestErr("password not correct", nil)
		c.JSON(loginErr.StatusCode, loginErr)
		return
	}

	if err := uc.userService.DeleteUser(user); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	response := ports.DeleteUserReply()

	c.JSON(http.StatusOK, response)
}

func GetCurrentUser(c *gin.Context) (*dtos.User, bool) {
	if c.Writer.Status() == http.StatusUnauthorized {
		err := errs.NewUnauthorizedErr("unauthorized access", c.Errors[0])
		c.JSON(http.StatusUnauthorized, err)
		return &dtos.User{}, false
	}

	usr, ok := c.Keys["user"]
	if !ok {
		err := errs.NewUnauthorizedErr("unauthorized access", errors.New("could not get user from session"))
		c.JSON(http.StatusUnauthorized, err)
		return &dtos.User{}, false
	}

	user := usr.(dtos.User)
	return &user, true
}
