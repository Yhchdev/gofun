package controller

import (
	"github.com/gin-gonic/gin"
	"gofun/service/user"
)

type userController struct {
	Controller
	userService user.Service
}

var UserController = &userController{
	userService: user.New(),
}

func (u *userController) Register(ctx *gin.Context) {

}

func (u *userController) Login(ctx *gin.Context) {

}
