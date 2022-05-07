package routes

import (
	"github.com/Zaida-3dO/goblin/adapters/driver/rest/controllers"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var authController = controllers.NewAuthController("")

func InitRouter() {
	router = gin.Default()

	mapAuthUrls()

	router.Run("127.0.0.1:5000")
}
