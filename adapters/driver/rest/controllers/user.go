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
}

type userController struct {
	userService services.UserService
}

func NewUserController(mode string) UserController {
	var uc UserController = &userController{
		userService: services.NewUserService(mode),
	}
	return uc
}

func (uc *userController) ChangePassword(c *gin.Context) {
	user, b := GetUserFromContext(c)
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

	c.JSON(http.StatusCreated, response)
}

func GetUserFromContext(c *gin.Context) (*dtos.User, bool) {
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
