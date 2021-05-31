package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//AuthContoller inteface is a contract whtat this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register (ctx *gin.Context)
}

type authController struct {

}

func NewAuthController() AuthController{
	return &authController{}
}

func (a authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,gin.H{
		"message" :"hello login",
	})
}

func (a authController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK,gin.H{
		"message" :"hello register",
	})
}
