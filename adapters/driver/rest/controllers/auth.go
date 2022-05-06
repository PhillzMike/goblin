package controllers

import (
	"net/http"

	"github.com/Zaida-3dO/goblin/internal/ports"
	"github.com/Zaida-3dO/goblin/internal/services"
	"github.com/Zaida-3dO/goblin/pkg/errs"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type authController struct {
	authService services.AuthService
}

func NewAuthController() AuthController {
	var ac AuthController = &authController{
		authService: services.NewAuthService(),
	}
	return ac
}

func (ac *authController) Register (c *gin.Context) {
	var request ports.RegisterUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		httpErr := errs.NewBadRequestErr("invalid json body", err)
		c.JSON(httpErr.StatusCode, httpErr)
		return
	}

	user, at, rt, saveErr := ac.authService.RegisterUser(&request)
	if saveErr != nil {
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}

	response := ports.NewLoginReply(user, at, rt, "account created successfully")

	c.JSON(http.StatusCreated, response)
}

func (ac *authController) Login(c *gin.Context) {
	var request ports.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		httpErr := errs.NewBadRequestErr("invalid json body", nil)
		c.JSON(httpErr.StatusCode, httpErr)
		return
	}

	user, at, rt, saveErr := ac.authService.LoginUser(&request)
	if saveErr != nil {
		c.JSON(saveErr.StatusCode, saveErr)
		return
	}

	response := ports.NewLoginReply(user, at, rt, "logged in successfully")

	c.JSON(http.StatusOK, response)
}