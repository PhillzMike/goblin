package routes

func mapAuthUrls() {
	router.POST("/v1/auth/register", authController.Register)
	router.POST("/v1/auth/login", authController.Login)
}
