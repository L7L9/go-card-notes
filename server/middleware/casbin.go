package middleware

import (
	"github.com/gin-gonic/gin"
	"lqlzzz/go-card-notes/global"
	"lqlzzz/go-card-notes/model/common/response"
	"lqlzzz/go-card-notes/service"
	"lqlzzz/go-card-notes/utils"
	"strconv"
	"strings"
)

var casbinService = service.ServiceOuter.CasbinService

// CasbinAuth //
// 权限控制中间件
func CasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取用户roleID
		roleID := utils.GetUserRoleID(c)
		// 初始化sub,obj,act
		sub := strconv.Itoa(int(roleID))
		withPrefixObj := c.Request.URL.Path
		// 处理路径
		obj := strings.TrimPrefix(withPrefixObj, global.GCN_CONFIG.System.RouterPrefix)
		act := c.Request.Method

		// 获取鉴权的casbin
		enforcer := casbinService.GetCasbin()

		// 鉴权
		success, _ := enforcer.Enforce(sub, obj, act)
		if !success {
			response.FailedWithMsg(c, "权限不足")
			c.Abort()
			return
		}
		c.Next()
	}
}
