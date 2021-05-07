package controller

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gofun/service/user"
)

// Controller:请求公共的返回结构体
// userService:service接口层
type userController struct {
	Controller
	userService user.Service
}

// UserController实例
var UserController = &userController{
	userService: user.New(),
}

func (u *userController) Register(ctx *gin.Context) {

	// 参数获取和校验,非必要(参数校验可放在前端)
	name := ctx.PostForm("name")
	if !govalidator.StringLength(name, "6", "10") {
		u.Failed(ctx, ParamError, "用户名长度不正确")
		return
	}

	email := ctx.PostForm("email")
	if !govalidator.IsEmail(email) {
		u.Failed(ctx, ParamError, "邮箱不正确")
	}

	password := ctx.PostForm("password")
	if !govalidator.StringLength(password, "6", "12") {
		u.Failed(ctx, ParamError, "密码长度不正确")
		return
	}

	// 调用service层注册
	token, err := u.userService.Register(name, password, email)
	if err != nil {
		if _, ok := err.(user.Error); ok {
			u.Failed(ctx, ParamError, err.Error())
		} else {
			u.Failed(ctx, Failed, "注册失败")
		}
	} else {
		u.Success(ctx, "ok", gin.H{"token": token})
	}

}

func (u *userController) Login(ctx *gin.Context) {
	name := ctx.PostForm("name")
	password := ctx.PostForm("password")

	token, err := u.userService.Login(name, password)

	if err != nil {
		if _, ok := err.(user.Error); ok {
			u.Failed(ctx, ParamError, err.Error())
		} else {
			u.Failed(ctx, Failed, "登录失败")
		}
	} else {
		u.Success(ctx, "ok", gin.H{"token": token})
	}
}
