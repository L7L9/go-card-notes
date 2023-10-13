package service

import (
	"errors"
	"gorm.io/gorm"
	"lqlzzz/go-card-notes/global"
	"lqlzzz/go-card-notes/model"
	"lqlzzz/go-card-notes/utils"
)

type UserService struct{}

// ChangePassword //
// 修改密码
func (service *UserService) ChangePassword(user *model.User, newPassword string) error {
	// 查询用户
	var tempUser model.User
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
func (service *UserService) UpdateUserInformation(user *model.User) (*model.User, error) {
	if err := global.GCN_DB.Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// OperateFollow //
// 关注or取消关注
func (service *UserService) OperateFollow(userFollow *model.UserFollow) error {
	// 开启事务
	tx := global.GCN_DB.Begin()
	defer tx.Commit()
	// 查询中间表
	var temp model.UserFollow
	if errors.Is(tx.Where("user_id = ? AND follow_id =?", userFollow.UserID, userFollow.FollowID).First(&temp).Error, gorm.ErrRecordNotFound) {
		// 查询不到记录，说明还没关注
		userFollow.Status = true
		tx.Create(userFollow)
	} else {
		// 查询到记录，检测status
		tx.Model(&temp).Update("status", !temp.Status)
	}

	// 操作用户表
	var tempUser1 model.User
	var tempUser2 model.User
	if userFollow.Status {
		// 用户关注数+1
		tx.First(&tempUser1, userFollow.UserID)
		tempUser1.FollowCount = tempUser1.FollowCount + 1
		tx.Save(tempUser1)
		// 被关注用户粉丝数+1
		tx.First(&tempUser2, userFollow.FollowID)
		tempUser2.FollowedCount = tempUser2.FollowedCount + 1
		tx.Save(tempUser2)
	} else {
		// 用户关注数-1
		tx.First(&tempUser1, userFollow.UserID)
		tempUser1.FollowCount = tempUser1.FollowCount - 1
		tx.Save(tempUser1)
		// 被关注用户粉丝数-1
		tx.First(&tempUser2, userFollow.FollowID)
		tempUser2.FollowedCount = tempUser2.FollowedCount - 1
		tx.Save(tempUser2)
	}

	return nil
}

// GetFollowerList //
// 获取粉丝
func (service *UserService) GetFollowerList(userID uint) ([]model.User, error) {
	// 定义管道
	var userFollow []model.UserFollow
	err := global.GCN_DB.Where("follow_id = ?", userID).Find(&userFollow).Error
	if err != nil {
		return nil, errors.New("数据库层面出错")
	}
	var result []model.User
	for _, v := range userFollow {
		var user model.User
		global.GCN_DB.Find(&user, v.UserID)
		result = append(result, user)
	}
	return result, nil
}

// GetFollowList //
// 获取粉丝
func (service *UserService) GetFollowList(userID uint) ([]model.User, error) {
	// 定义管道
	var userFollow []model.UserFollow
	err := global.GCN_DB.Where("user_id = ?", userID).Find(&userFollow).Error
	if err != nil {
		return nil, errors.New("数据库层面出错")
	}
	var result []model.User
	for _, v := range userFollow {
		var user model.User
		global.GCN_DB.Find(&user, v.FollowID)
		result = append(result, user)
	}
	return result, nil
}

// GetUserById //
// 通过id获取用户信息
func (service *UserService) GetUserById(id uint) (*model.User, error) {
	var user model.User
	if err := global.GCN_DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
