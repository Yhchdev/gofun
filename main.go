package main

import (
	"gofun/common"
	_ "gofun/common"
	_ "gofun/jwt"
	log "gofun/log"
	_ "gofun/oss"
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
}

func main() {
	log.Info(common.GlobalConfig.OSS.Bucket)
}
