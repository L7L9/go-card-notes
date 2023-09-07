package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"lqlzzz/go-card-notes/global"
	"time"
)

type Jwt struct {
	SigningKey []byte
}

// SystemClaims //
// 存于token令牌中的数据
type SystemClaims struct {
	UserID uint
	RoleID uint
	jwt.RegisteredClaims
}

// NewJwt //
// 获取一个jwt
func NewJwt() *Jwt {
	return &Jwt{SigningKey: []byte(global.GCN_CONFIG.Jwt.SigningKey)}
}

// GenerateClaims //
// 生成Claims
func (j *Jwt) GenerateClaims(userId, roleId uint) *SystemClaims {
	// 计算超时时间
	duration := time.Second * time.Duration(global.GCN_CONFIG.Jwt.ExpiresTime)
	return &SystemClaims{
		UserID: userId,
		RoleID: roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.GCN_CONFIG.Jwt.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
}

// GenerateToken //
// 创建令牌
func (j *Jwt) GenerateToken(claims *SystemClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SigningString()
}

// ParseToken //
// 解析token
func (j *Jwt) ParseToken(tokenString string) (*SystemClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &SystemClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	// 判断有无错误
	if err != nil {
		return nil, err
	}
	if token != nil {
		// 类型断言
		if claim, ok := token.Claims.(*SystemClaims); ok && token.Valid {
			return claim, nil
		} else {
			return nil, errors.New("token invalid")
		}
	}
	return nil, errors.New("token invalid")
}
