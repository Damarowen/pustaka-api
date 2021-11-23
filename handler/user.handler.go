package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/JWT"
	"pustaka-api/dto"
	"pustaka-api/helper"
	"pustaka-api/user"

	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//UserController is a ....
type IUserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type UserController struct {
	userService user.IUserService
	jwtService  JWT.IJwtService
}

//NewUserController is creating anew instance of UserControlller
func NewUserController(userServ user.IUserService, jwtService JWT.IJwtService) IUserController {
	return &UserController{
		userService: userServ,
		jwtService:  jwtService,
	}
}

func (c *UserController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u, ok , err:= c.userService.Update(userUpdateDTO)

	if !ok  {
		res := helper.BuildErrorResponse("Failed to process request",err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *UserController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}