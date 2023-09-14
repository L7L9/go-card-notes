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
func (service *BaseService) SignUp(user *schema.User) error {
	var tempUser schema.User
	// 判断有无相同账号名的用户
	if !errors.Is(global.GCN_DB.Where("username = ?", user.Username).First(&tempUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("账号名字重复，请重新输入")
	}

	// 加密password
	user.Password = utils.HashEncrypt(user.Password)

	return global.GCN_DB.Create(&user).Error
}

// SignIn //
// 登录
func (service *BaseService) SignIn(user *schema.User) (*schema.User, error) {
	// 处理password
	user.Password = utils.HashEncrypt(user.Password)

	var loginUser schema.User
	// 查找有无该用户
	if err := global.GCN_DB.Where("username = ?", user.Username).First(&loginUser).Error; err != nil {
		return nil, errors.New("账号查询不到记录，请输入正确的账号")
	}
	// 判断密码是否正确
	if ok := utils.CompanyHash(user.Password, loginUser.Password); !ok {
		return nil, errors.New("密码错误，请重新输入")
	}

	return &loginUser, nil
}
