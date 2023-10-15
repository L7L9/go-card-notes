package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// User //
// gcn_user table
type User struct {
	gorm.Model
	UUID          uuid.UUID   `json:"uuid" gorm:"NOT NULL;index;comment:用户UUID"`
	Username      string      `json:"username" gorm:"NOT NULL;comment:用户名"`
	Password      string      `json:"-" gorm:"NOT NULL;comment:用户密码"`
	Nickname      string      `json:"nickname" gorm:"NOT NULL;comment:昵称"`
	HeadImg       string      `json:"headImg" gorm:"default:'';NOT NULL;comment:头像图片"`
	Phone         string      `json:"phone" gorm:"default:'';NOT NULL;comment:手机号"`
	Email         string      `json:"email" gorm:"default:'';NOT NULL;comment:邮箱"`
	Points        int         `json:"points" gorm:"default:0;NOT NULL;comment:用户积分"`
	NoteGroup     []NoteGroup `json:"noteGroup" gorm:""`
	NoteCount     int         `json:"noteCount" gorm:"default:0;NOT NULL;comment:用户笔记数"`
	FollowCount   int         `json:"followCount" gorm:"default:0;NOT NULL;comment:关注数"`
	FollowedCount int         `json:"followedCount" gorm:"default:0;NOT NULL;comment:粉丝数"`
	Introduction  string      `json:"introduction" gorm:"default:'';NOT NULL;comment:简介"`
	BaseRoleID    uint        `json:"baseRoleID" gorm:"default:3;NOT NULL;comment:角色id"`
	Roles         []Role      `json:"roles" gorm:"many2many:gcn_user_role;"`
	Status        int         `json:"status" gorm:"default:1;NOT NULL;comment:用户状态:1=>正常;2=>注销;3=>冻结"`
}

func (User) TableName() string {
	return "gcn_user"
}
