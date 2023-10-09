package initialize

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"lqlzzz/go-card-notes/global"
	"lqlzzz/go-card-notes/model"
)

// InitDbTable //
// 初始化数据库中的表
func InitDbTable() {
	err := global.GCN_DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.Note{},
		&model.Tag{},
	)

	if err != nil {
		global.GCN_LOGGER.Error("initialize schema error: ", zap.Error(err))
	}
	if err = initPolicy(); err != nil {
		global.GCN_LOGGER.Error("initialize policy table error: ", zap.Error(err))
	}
	return
}

// initPolicy //
// 初始化策略表(权限表)
func initPolicy() error {
	// 定义策略
	// V0: 1=>超级管理员 2=>管理员 3=>普通用户
	entities := []gormadapter.CasbinRule{
		// 定义超级管理员的权限
		{Ptype: "p", V0: "1", V1: "/v1/changePassword", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/v1/updateUserInformation", V2: "PUT"},
		{Ptype: "p", V0: "1", V1: "/v1/operateFollow", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/v1/getFollowList", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/v1/getFollowerList", V2: "GET"},

		// 定义管理员的权限
		{Ptype: "p", V0: "2", V1: "/v1/changePassword", V2: "PUT"},
		{Ptype: "p", V0: "2", V1: "/v1/updateUserInformation", V2: "PUT"},
		{Ptype: "p", V0: "2", V1: "/v1/operateFollow", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/v1/getFollowList", V2: "GET"},
		{Ptype: "p", V0: "2", V1: "/v1/getFollowerList", V2: "GET"},

		// 定义普通用户的权限
		{Ptype: "p", V0: "3", V1: "/v1/changePassword", V2: "PUT"},
		{Ptype: "p", V0: "3", V1: "/v1/updateUserInformation", V2: "PUT"},
		{Ptype: "p", V0: "3", V1: "/v1/operateFollow", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/v1/getFollowList", V2: "GET"},
		{Ptype: "p", V0: "3", V1: "/v1/getFollowerList", V2: "GET"},
	}

	if err := global.GCN_DB.AutoMigrate(&gormadapter.CasbinRule{}); err != nil {
		return err
	}
	return global.GCN_DB.Create(&entities).Error
}
