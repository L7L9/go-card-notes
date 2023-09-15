package utils

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"lqlzzz/go-card-notes/global"
)

// GetUserInfo //
// 通过请求头获取claims
func GetUserInfo(c *gin.Context) (*SystemClaims, error) {
	token := c.Request.Header.Get("Token")
	j := NewJwt()
	if claims, err := j.ParseToken(token); err != nil {
		global.GCN_LOGGER.Error("获取claims失败，请检查token或者claims格式")
		return nil, err
	} else {
		return claims, nil
	}
}

// GetUserID //
// 获取用户id
func GetUserID(c *gin.Context) uint {
	if claims, isExist := c.Get("claims"); !isExist {
		if cl, err := GetUserInfo(c); err != nil {
			return 0
		} else {
			return cl.UserID
		}
	} else {
		userInfo := claims.(*SystemClaims)
		return userInfo.UserID
	}
}

// GetUserUUID //
// 获取用户uuid
func GetUserUUID(c *gin.Context) uuid.UUID {
	if claims, isExist := c.Get("claims"); !isExist {
		if cl, err := GetUserInfo(c); err != nil {
			return uuid.UUID{}
		} else {
			return cl.UUID
		}
	} else {
		userInfo := claims.(*SystemClaims)
		return userInfo.UUID
	}
}

// GetUserRoleID //
// 获取用户角色id
func GetUserRoleID(c *gin.Context) uint {
	if claims, isExist := c.Get("claims"); !isExist {
		if cl, err := GetUserInfo(c); err != nil {
			return 0
		} else {
			return cl.RoleID
		}
	} else {
		userInfo := claims.(*SystemClaims)
		return userInfo.RoleID
	}
}
