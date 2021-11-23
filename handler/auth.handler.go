package handler

import (
	"net/http"
	"pustaka-api/JWT"
	"pustaka-api/auth"
	"pustaka-api/dto"
	"pustaka-api/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type IAuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type AuthController struct {
	authService auth.IAuthService
	jwtService  JWT.IJwtService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService auth.IAuthService, jwtService JWT.IJwtService) IAuthController {
	return &AuthController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)

	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult, err := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)

	if err != nil {
		response := helper.BuildErrorResponse("Please check again your credential",err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	//* logic dibawah balikanya boolean
	user := authResult

	generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(user.ID, 10))
	user.Token = generatedToken
	response := helper.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, response)

	// response := helper.BuildErrorResponse("Please check again your credential", "Invalid PASSWORD", helper.EmptyObj{})
	// ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *AuthController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
