package main

import (
	"AP/golang-api-jwt/config"
	"AP/golang-api-jwt/controller"
	"github.com/gin-gonic/gin"
)
var (
	db = config.SetupDatabaseConnection()
	authController = controller.NewAuthController()

)
func main()  {
	router := gin.Default()
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/login",authController.Login)
		authRoutes.POST("/register",authController.Register)
	}
	router.Run()
}
