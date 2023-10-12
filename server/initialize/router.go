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

	// 不做鉴权的路由组
	publicGroup := r.Group(global.GCN_CONFIG.System.RouterPrefix)
	{
		// 初始化基础路由
		router.RouterGroupOuter.BaseRouter.Initialize(publicGroup)
	}

	// 要做鉴权的路由组
	privateGroup := r.Group(global.GCN_CONFIG.System.RouterPrefix, middleware.JwtAuth(), middleware.CasbinAuth())
	{
		// 初始化路由
		router.RouterGroupOuter.UserRouter.Initialize(privateGroup)
	}

	return r
}
