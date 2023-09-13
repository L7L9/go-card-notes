package middleware

import (
	"github.com/gin-gonic/gin"
	"lqlzzz/go-card-notes/model/common/response"
	"lqlzzz/go-card-notes/utils"
)

// JwtAuth //
// 验证token的中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的token
		token := c.Request.Header.Get("Token")

		// 判断有无token
		if token == "" {
			response.FailedWithMsg(c, "没有携带token，请登录")
			// Abort方法会终止请求链中的后续处理
			c.Abort()
			return
		}

		// TODO 判断token有无过期
		j := utils.NewJwt()
		claims, err := j.ParseToken(token)
		if err != nil {
			response.FailedWithMsg(c, "token无效")
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
