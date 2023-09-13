package service

import (
	"github.com/casbin/casbin/v2"
	model2 "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"lqlzzz/go-card-notes/global"
	"sync"
)

type CasbinService struct {
}

var (
	enforcer casbin.CachedEnforcer
	once     sync.Once
)

// GetCasbin //
// 获取Casbin验权的结构体
func (service *CasbinService) GetCasbin() casbin.CachedEnforcer {
	// 使用once懒加载,
	once.Do(func() {
		// 加载gorm设配器
		adaptor, err := gormadapter.NewAdapterByDB(global.GCN_DB)
		if err != nil {
			global.GCN_LOGGER.Error("初始化casbin-gorm-adaptor失败，请检查响应结构体")
			return
		}

		// 设置casbin中的model
		modelText := `
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act && keyMatch2(r.obj,p.obj)
		`
		// 加载模型
		model, err := model2.NewModelFromString(modelText)
		if err != nil {
			global.GCN_LOGGER.Error("casbin-model加载失败，请检查modelText", zap.Error(err))
			return
		}
		// 初始化enforcer
		enforcer, _ := casbin.NewCachedEnforcer(model, adaptor)
		// 加载策略Policy
		_ = enforcer.LoadPolicy()
	})
	return enforcer
}
