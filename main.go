package main

import (
	"github.com/gin-gonic/gin"
	"gofun/common"
	"gofun/database/mysql"
	log "gofun/log"
	_ "gofun/oss"
	"gofun/router"
	_ "gofun/session"
	_ "gofun/socket"
)

func init() {
	// 初始化配置文件
	err := common.InitConfig("./conf/conf.yaml")
	if err != nil {
		log.Error(err)
	}

	//初始化log配置
	log.InitLog()

	//初始化mysql
	mysql.Init()
}

func main() {
	log.Info(common.GlobalConfig.OSS.Bucket)
	// todo：gin配置写入log中
	gin.SetMode(gin.DebugMode)
	engine := gin.New()

	router.Router(engine)

	engine.Run(":9999")
}
