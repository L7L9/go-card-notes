package model

import "gorm.io/gorm"

// Note //
// gcn_note table
type Note struct {
	gorm.Model
	Title        string `json:"title" gorm:"NOT NULL;comment:笔记标题"`
	AuthorName   string `json:"authorName" gorm:"NOT NULL;comment:作者昵称"`
	Uri          string `json:"uri" gorm:"NOT NULL;comment:笔记路径"`
	CollectCount int    `json:"collectCount" gorm:"NOT NULL;default:0;comment:笔记标题"`
	GroupID      string `json:"groupId" gorm:"NOT NULL;comment:分组id"`
	Status       int    `json:"status" gorm:"NOT NULL;comment:笔记状态(0=>公开;1=>私藏)"`
}

func (Note) TableName() string {
	return "gcn_note"
}
