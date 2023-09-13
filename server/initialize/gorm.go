package initialize

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lqlzzz/go-card-notes/global"
)

// InitGorm //
// initialize mysql use gorm
func InitGorm() *gorm.DB {
	mysqlConfig := global.GCN_CONFIG.Mysql

	db, err := gorm.Open(mysql.Open(mysqlConfig.Dsn()), &gorm.Config{})
	if err != nil {
		global.GCN_LOGGER.Error("failed to connect database: ", zap.Error(err))
		return nil
	}
	sqlDB, _ := db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConn)
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConn)
	return db
}
