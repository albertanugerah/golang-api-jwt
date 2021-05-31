package controller

import (
	"AP/golang-api-jwt/dto"
	"AP/golang-api-jwt/entity"
	"AP/golang-api-jwt/helper"
	"AP/golang-api-jwt/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//AuthContoller inteface is a contract whtat this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register (ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController{
	return &authController{
		authService : authService,
		jwtService : jwtService,
	}
}

func (a authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failded to process resquest",errDTO.Error(),helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
	authResult := a.authService.VerifyCredential(loginDTO.Email,loginDTO.Password)
	if value, ok := authResult.(entity.User); ok{
		generateToken := a.jwtService.GenerateToken(strconv.FormatUint(value.ID, 10))
		value.Token = generateToken
		response := helper.BuildResponse(true,"OK", value)
		ctx.JSON(http.StatusOK,response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential","Invalid Credential",helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized,response)
}

func (a authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil{
		response := helper.BuildErrorResponse("Failed to process request",errDTO.Error(),helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	if !a.authService.IsDuplicateEmail(registerDTO.Email){
		response := helper.BuildErrorResponse("Failed to process request","duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict,response)
	}else{
		createdUser := a.authService.CreateUser(registerDTO)
		token := a.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID,10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK",createdUser)
		ctx.JSON(http.StatusCreated,response)
	}
}
