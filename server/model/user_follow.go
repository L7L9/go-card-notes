package model

import "gorm.io/gorm"

// UserFollow //
// 用户关注的中间表
type UserFollow struct {
	gorm.Model
	UserID   uint `gorm:"NOT NULL;column:user_id;comment:用户id"`
	FollowID uint `gorm:"NOT NULL;column:follow_id;comment:关注的用户id"`
	Status   bool `gorm:"NOT NULL;comment:关注的状态"`
}

func (UserFollow) TableName() string {
	return "gcn_user_follow"
}
