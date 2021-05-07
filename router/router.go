package router

import (
	"github.com/gin-gonic/gin"
	. "gofun/controller"
)

func Router(router *gin.Engine) {
	//todo :使用中间件
	router.Use()

	// 路由分组
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.Group("/user").
				POST("/register", UserController.Register).
				POST("/login", UserController.Login)
		}
	}
}
