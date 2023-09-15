package main

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"lqlzzz/go-card-notes/global"
	"lqlzzz/go-card-notes/initialize"
)

func main() {
	// 初始配置
	global.GCN_VIPER = initialize.InitViper()
	// 初始日志类
	global.GCN_LOGGER = initialize.InitZap()
	// 初始化数据库
	global.GCN_DB = initialize.InitGorm()
	// 初始化路由
	router := initialize.InitRouter()
	if global.GCN_DB != nil {
		// 初始化表
		initialize.InitDbTable()
		// 程序结束前关闭数据库链接
		db, _ := global.GCN_DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				global.GCN_LOGGER.Error("数据库连接关闭失败", zap.Error(err))
			}
		}(db)
	}
	// TODO
	port := fmt.Sprintf(":%d", global.GCN_CONFIG.System.Addr)
	if err := router.Run(port); err != nil {
		global.GCN_LOGGER.Error("服务启动失败", zap.Error(err))
	}
}
