package model

import "gorm.io/gorm"

// NoteGroup //
// 笔记分组
type NoteGroup struct {
	gorm.Model
	CreatorID   uint   `gorm:"NOT NULL;commit:分组创建者id"`
	Name        string `gorm:"NOT NULL;commit:笔记组名"`
	Description string `gorm:"NOT NULL;default:'';commit:分组描述"`
	Note        []Note
}

func (NoteGroup) TableName() string {
	return "gcn_note_group"
}
