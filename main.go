package main

import (
	"github.com/albertanugerah/golang-api-jwt/config"
	"github.com/albertanugerah/golang-api-jwt/controller"
	"github.com/albertanugerah/golang-api-jwt/middleware"
	"github.com/albertanugerah/golang-api-jwt/repository"
	"github.com/albertanugerah/golang-api-jwt/service"
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
