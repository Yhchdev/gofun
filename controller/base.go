package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

func (*Controller) Success(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    code,
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
