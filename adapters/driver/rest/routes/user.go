package routes

import (
	"github.com/Zaida-3dO/goblin/adapters/driver/rest/middlewares"
)

func mapUserUrls() {
	userRouter := router.Group("/v1/users", middlewares.Authorization())

	userRouter.PUT("/change-password", userController.ChangePassword)
	userRouter.PUT("/", userController.UpdateUser)
}
