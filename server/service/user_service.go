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
	var temp model.User
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

func (service *UserService) Follow(userFollow *model.UserFollow) error {
	// 开启事务
	tx := global.GCN_DB.Begin()
	defer tx.Commit()
	// 查询中间表
	var temp model.UserFollow
	if errors.Is(tx.Where("user_id = ? AND follow_id =?", userFollow.UserID, userFollow.FollowID).First(&temp).Error, gorm.ErrRecordNotFound) {
		// 查询不到记录，说明还没关注
		if !userFollow.Status {
			return errors.New("前端参数输入错误")
		}
		// 插入数据
		tx.Create(userFollow)
	} else {
		// 查询到记录，检测status
		if temp.Status != userFollow.Status {
			tx.Model(&temp).Update("status", userFollow.Status)
			// 操作用户表
			var tempUser model.User
			if userFollow.Status {
				// 用户关注数+1
				tx.First(&tempUser, userFollow.UserID)
				tempUser.FollowCount = tempUser.FollowCount + 1
				tx.Save(tempUser)
				// 被关注用户粉丝数+1
				tx.First(&tempUser, userFollow.FollowID)
				tempUser.FollowCount = tempUser.FollowedCount + 1
				tx.Save(tempUser)
			} else {
				// 用户关注数-1
				tx.First(&tempUser, userFollow.UserID)
				tempUser.FollowCount = tempUser.FollowCount - 1
				tx.Save(tempUser)
				// 被关注用户粉丝数-1
				tx.First(&tempUser, userFollow.FollowID)
				tempUser.FollowCount = tempUser.FollowedCount - 1
				tx.Save(tempUser)
			}
		} else {
			return errors.New("前端参数输入错误")
		}
	}
	return nil
}
