package main

import (
	"gofun/common"
	_ "gofun/common"
	_ "gofun/jwt"
	_ "gofun/oss"
	_ "gofun/session"
	"log"
)

func main() {
	err := common.InitConfig("./conf/conf.yaml")
	if err != nil {
		log.Println(err)
	}

	log.Println(common.GlobalConfig.OSS.Bucket)
}
