package routes

import (
	"github.com/Zaida-3dO/goblin/adapters/driver/rest/controllers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var authController controllers.AuthController
var userController controllers.UserController

func Init() {
	authController = controllers.NewAuthController("psql")
	userController = controllers.NewUserController("psql")

	router = gin.Default()

	mapAuthUrls()
	mapUserUrls()

	router.Run("127.0.0.1:5000")
}
