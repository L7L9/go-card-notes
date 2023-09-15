package service

import (
	"errors"
	"lqlzzz/go-card-notes/global"
	"lqlzzz/go-card-notes/model/schema"
	"lqlzzz/go-card-notes/utils"
)

type UserService struct{}

// ChangePassword //
// 修改密码
func (service *UserService) ChangePassword(user *schema.User, newPassword string) error {
	// 查询用户
	var tempUser schema.User
	if err := global.GCN_DB.Where("id = ?", user.ID).First(&tempUser).Error; err != nil {
		return err
	}
	// 验证原密码是否正确
	if !utils.CompanyHash(user.Password, tempUser.Password) {
		return errors.New("原密码不符合，请重新输入")
	}
	// 验证新旧密码是否一致
	if utils.CompanyHash(newPassword, tempUser.Password) {
		return errors.New("新密码与旧密码一样，请重新输入")
	}
	tempUser.Password = utils.HashEncrypt(newPassword)
	return global.GCN_DB.Save(&tempUser).Error
}

// UpdateUserInformation //
// 修改用户信息
func (service *UserService) UpdateUserInformation(user *schema.User) (*schema.User, error) {
	var temp schema.User
	err := global.GCN_DB.Where("id = ?", user.ID).First(user).Error
	if err != nil {
		return nil, err
	}
	temp.Nickname = user.Nickname
	temp.Email = user.Email
	temp.Phone = user.Phone
	temp.HeadImg = user.HeadImg
	if err = global.GCN_DB.Save(&temp).Error; err != nil {
		return nil, err
	}
	return &temp, nil
}
