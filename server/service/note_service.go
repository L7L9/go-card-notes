package service

import (
	"errors"
	"gorm.io/gorm"
	"lqlzzz/go-card-notes/global"
	"lqlzzz/go-card-notes/model"
)

type NoteService struct {
}

// CreateNote //
// 创建笔记
func (service *NoteService) CreateNote() error {
	return nil
}

// CreateNoteGroup //
// 创建笔记分组
func (service *NoteService) CreateNoteGroup(group *model.NoteGroup) error {
	var tempGroup model.NoteGroup
	if errors.Is(global.GCN_DB.Where("creator_id =? AND name = ?", group.CreatorID, group.Name).First(&tempGroup).Error, gorm.ErrRecordNotFound) {
		return errors.New("无法创建同名分组")
	}

	return global.GCN_DB.Create(&group).Error
}
