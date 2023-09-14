package initialize

import (
	"go.uber.org/zap"
	"lqlzzz/go-card-notes/global"
	"lqlzzz/go-card-notes/model/schema"
)

// InitDbTable //
// 初始化数据库中的表
func InitDbTable() {
	err := global.GCN_DB.AutoMigrate(
		&schema.User{},
		&schema.Role{},
	)

	if err != nil {
		global.GCN_LOGGER.Error("initialize schema error: ", zap.Error(err))
	}
	return
}
