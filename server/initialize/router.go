package initialize

import (
	"github.com/gin-gonic/gin"
	"lqlzzz/go-card-notes/global"
	"lqlzzz/go-card-notes/middleware"
	"lqlzzz/go-card-notes/router"
)

// InitRouter //
// initialize router and group
func InitRouter() *gin.Engine {
	r := gin.Default()
	// 设置前缀
	prefixGroup := r.Group(global.GCN_CONFIG.System.RouterPrefix)

	// 不做鉴权的路由组
	publicGroup := prefixGroup.Group("/v1")
	{
		// 初始化基础路由
		router.RouterGroupOuter.BaseRouter.Initialize(publicGroup)
	}

	// 要做鉴权的路由组
	privateGroup := prefixGroup.Group("/v1", middleware.JwtAuth(), middleware.CasbinAuth())
	{
		// 初始化路由
		router.RouterGroupOuter.UserRouter.Initialize(privateGroup)
	}

	return r
}
