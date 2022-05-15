package routes

func mapAuthUrls() {
	router.POST("/v1/auth/register", authController.Register)
	router.POST("/v1/auth/login", authController.Login)
	router.POST("/v1/auth/forgot-password", authController.ForgotPassword)
	router.POST("/v1/auth/reset-password", authController.ResetPassword)
}
