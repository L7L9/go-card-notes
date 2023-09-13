package schema

import (
	"gorm.io/gorm"
)

// User //
// gcn_user table
type User struct {
	gorm.Model
	Username      string `gorm:"NOT NULL;comment:'用户名'"`
	Password      string `gorm:"NOT NULL;comment:'用户密码'"`
	Nickname      string `gorm:"NOT NULL;comment:'昵称'"`
	HeadImg       string `gorm:"default:0;NOT NULL;comment:'头像图片'"` // TODO
	Phone         string `gorm:"default:'';NOT NULL;comment:'手机号'"`
	Email         string `gorm:"default:'';NOT NULL;comment:'邮箱'"`
	Points        int    `gorm:"default:0;NOT NULL;comment:'用户积分'"`
	FollowCount   int    `gorm:"default:0;NOT NULL;comment:'关注数'"`
	FollowedCount int    `gorm:"default:0;NOT NULL;comment:'粉丝数'"`
	Introduction  string `gorm:"default:'';NOT NULL;comment:'简介'"`
	BaseRoleID    uint   `gorm:"default:3;NOT NULL;comment:'角色id'"`
	Roles         []Role `gorm:"many2many:'gcn_user_role';"`
	Status        int    `gorm:"default:1;NOT NULL;comment:'用户状态:1=>正常;2=>注销;3=>冻结'"`
}

func (User) TableName() string {
	return "gcn_user"
}
