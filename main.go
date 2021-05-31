package main

import (
	"AP/golang-api-jwt/config"
	"AP/golang-api-jwt/controller"
	"AP/golang-api-jwt/middleware"
	"AP/golang-api-jwt/repository"
	"AP/golang-api-jwt/service"
	"github.com/gin-gonic/gin"
)
var (
	db = config.SetupDatabaseConnection()
	userRepository 		= repository.NewUserRepository(db)
	jwtService			= service.NewJWTService()
	userService			= service.NewUserService(userRepository)
	authService			= service.NewAuthService(userRepository)
	authController 		= controller.NewAuthController(authService,jwtService)
	userController		= controller.NewUserController(userService,jwtService)

)
func main()  {
	router := gin.Default()
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/login",authController.Login)
		authRoutes.POST("/register",authController.Register)
	}
	userRoutes := router.Group("api/user",middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile",userController.Profile)
		userRoutes.PUT("/profile",userController.Update)
	}
	router.Run()
}
