package initialize

import (
	"go.uber.org/zap"
	"lqlzzz/go-card-notes/global"
	"lqlzzz/go-card-notes/model"
)

// InitDbTable //
// 初始化数据库中的表
func InitDbTable() {
	err := global.GCN_DB.AutoMigrate(
		&model.User{},
		&model.UserFollow{},
		&model.Role{},
		&model.UserRole{},
		&model.Note{},
		&model.NoteGroup{},
	)

	if err != nil {
		global.GCN_LOGGER.Error("initialize schema error: ", zap.Error(err))
	}
	if err = initPolicy(); err != nil {
		global.GCN_LOGGER.Error("initialize policy table error: ", zap.Error(err))
	}
	return
}
