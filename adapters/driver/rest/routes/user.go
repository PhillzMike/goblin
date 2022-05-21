package routes

import (
	"github.com/Zaida-3dO/goblin/adapters/driver/rest/middlewares"
)

func mapUserUrls() {
	userRouter := router.Group("/v1/users", middlewares.Authorization())

	userRouter.POST("/change-password", userController.ChangePassword)
}
