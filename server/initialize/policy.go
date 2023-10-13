package initialize

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"lqlzzz/go-card-notes/global"
)

// initPolicy //
// 初始化策略表(权限表)
func initPolicy() error {
	// TODO 优化策略表的初始化
	// 定义策略
	// V0: 1=>超级管理员 2=>管理员 3=>普通用户
	entities := []gormadapter.CasbinRule{
		// 定义超级管理员的权限
		{Ptype: "p", V0: "1", V1: "/user/changePassword/", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/updateInformation/", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/follow/:id", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/getFollowList/", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/user/getFollowerList/", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/user/getUserInformation/:id", V2: "GET"},

		// 定义管理员的权限
		{Ptype: "p", V0: "2", V1: "/user/changePassword/", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/user/updateInformation/", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/user/follow/:id", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/user/getFollowList/", V2: "GET"},
		{Ptype: "p", V0: "2", V1: "/user/getFollowerList/", V2: "GET"},
		{Ptype: "p", V0: "2", V1: "/user/getUserInformation/:id", V2: "GET"},

		// 定义普通用户的权限
		{Ptype: "p", V0: "3", V1: "/user/changePassword/", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/user/updateInformation/", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/user/follow/:id", V2: "POST"},
		{Ptype: "p", V0: "3", V1: "/user/getFollowList/", V2: "GET"},
		{Ptype: "p", V0: "3", V1: "/user/getFollowerList/", V2: "GET"},
		{Ptype: "p", V0: "3", V1: "/user/getUserInformation/:id", V2: "GET"},
	}

	var err error
	if global.GCN_DB.Migrator().HasTable(&gormadapter.CasbinRule{}) {
		if err = global.GCN_DB.Migrator().DropTable(&gormadapter.CasbinRule{}); err != nil {
			return err
		}
	}
	if err = global.GCN_DB.AutoMigrate(&gormadapter.CasbinRule{}); err != nil {
		return err
	}
	return global.GCN_DB.Create(entities).Error
}
