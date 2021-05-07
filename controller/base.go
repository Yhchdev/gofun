package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success      = 200 //正常
	Failed       = 500 //失败
	ParamError   = 400 //参数错误
	NotFound     = 404 //不存在
	UnAuthorized = 401 //未授权
	NotLogin     = 405 //未登录
)

type Controller struct {
}

func (*Controller) Success(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    Success,
		"data":    data,
		"message": msg,
	})
}

func (c *Controller) Failed(ctx *gin.Context, code int, msg string) {
	ctx.Abort()
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
	})
}
