package model

import "gorm.io/gorm"

// Role //
// gcn_role table
type Role struct {
	gorm.Model
	Name string `gorm:"NOT NULL;comment:角色名字"`
}

func (Role) TableName() string {
	return "gcn_role"
}
