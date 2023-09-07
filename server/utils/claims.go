package utils

import (
	"github.com/gin-gonic/gin"
	"lqlzzz/go-card-notes/global"
)

// GetUserInfo //
// 获取context中的claims
func GetUserInfo(c *gin.Context) *SystemClaims {
	if claims, isExist := c.Get("claims"); !isExist {
		token := c.Request.Header.Get("Token")
		j := NewJwt()
		if cl, err := j.ParseToken(token); err != nil {
			global.GCN_LOGGER.Error("获取claims失败，请检查token或者claims格式")
			return nil
		} else {
			return cl
		}
	} else {
		userInfo := claims.(*SystemClaims)
		return userInfo
	}
}
