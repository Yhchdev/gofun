package mysql

import (
	"gofun/common"
	log "gofun/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func DBConn() *gorm.DB {
	return db
}

func Init() {
	gdb, err := gorm.Open(mysql.Open(common.GlobalConfig.DB.DSN), &gorm.Config{})
	if err != nil {
		log.Error("连接数据库失败:", err)
	}

	// 配置gorm连接池
	dbPool, err := gdb.DB()
	if err != nil {
		log.Error("连接数据库失败:", err)
	}
	dbPool.SetMaxOpenConns(common.GlobalConfig.DB.MaxOpen)
	dbPool.SetMaxIdleConns(common.GlobalConfig.DB.MaxIdle)

	db = gdb
}
