package service

import (
	"errors"
	"gorm.io/gorm"
	"lqlzzz/go-card-notes/global"
	"lqlzzz/go-card-notes/model/schema"
	"lqlzzz/go-card-notes/utils"
)

type BaseService struct{}

// SignUp //
// 注册
func (BaseService) SignUp(user schema.User) error {
	var tempUser schema.User
	// 判断有无相同账号名的用户
	if errors.Is(global.GCN_DB.Where("username = ?", user.Username).First(&tempUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("账号名字重复，请重新输入")
	}

	// 加密password
	user.Password = utils.HashEncrypt(user.Password)

	return global.GCN_DB.Create(&user).Error
}
