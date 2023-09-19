package model

import "gorm.io/gorm"

// Tag //
// 标签
type Tag struct {
	gorm.Model
	Name string `gorm:"NOT NULL;comment:'标签名'"`
}

func (Tag) TableName() string {
	return "gcn_tag"
}
