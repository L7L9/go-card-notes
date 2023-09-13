package middleware

import (
	"github.com/gin-gonic/gin"
	"lqlzzz/go-card-notes/model/common/response"
	"lqlzzz/go-card-notes/service"
	"lqlzzz/go-card-notes/utils"
	"strconv"
)

var casbinService = service.ServiceOuter.CasbinService

// CasbinAuth //
// 权限控制中间件
func CasbinAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 通过tokenID获取claims
		userClaims := utils.GetUserInfo(c)
		// 初始化sub,obj,act
		sub := strconv.Itoa(int(userClaims.RoleID))
		obj := c.Request.URL.Path
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
