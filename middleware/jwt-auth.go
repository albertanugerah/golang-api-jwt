package middleware

import (
	"github.com/albertanugerah/golang-api-jwt/helper"
	"github.com/albertanugerah/golang-api-jwt/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
//AuthorizeJWT validates token user given
func AuthorizeJWT(service service.JWTService) gin.HandlerFunc  {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == ""{
			response := helper.BuildErrorResponse("Failed to process request", "No Token Found",nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
		}
		token, err := service.ValidateToken(authHeader)
		if token.Valid{
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id] :",claims["user_id"])
			log.Println("Claim[issuer]: ",claims["issuer"])

		}else{
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(),nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized,response)
		}

	}

}
